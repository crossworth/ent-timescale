package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"ent-timescale/ent"
	"ent-timescale/ent/migrate"
	"ent-timescale/ent/sensor"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"github.com/rs/xid"
)

// https://docs.timescale.com/api/latest/hypertable/create_hypertable/

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=timescaledb password=root sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	ctx := context.Background()

	// automatic migrations
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		WithoutPrimaryKey(sensor.Table),
		EnableTimescaleDBOption(db),
		CreateHypertable(db, sensor.Table, sensor.FieldCreateTime),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// delete the data present
	client.Sensor.Delete().ExecX(ctx)

	// create a few fake sensors
	s1 := newTempSensor(-49.276855, -25.441105)
	s2 := newTempSensor(34.855499, 32.109333)
	s3 := newTempSensor(106.660172, 10.762622)

	// store temperature on each one of them
	recordTemperature(client, 1000, s1)
	recordTemperature(client, 1000, s2)
	recordTemperature(client, 1000, s3)

	// querying

	// number of temperature reading's by day
	var tempByDay []struct {
		Day   string
		Count int
	}
	client.Sensor.Query().
		GroupBy("1").
		Aggregate(func(selector *entsql.Selector) string {
			return "date_trunc('day', " + sensor.FieldCreateTime + ") AS day"
		}, ent.Count()).
		ScanX(ctx, &tempByDay)
	for _, r := range tempByDay {
		_ = r
		// fmt.Println(r.Day, r.Count)
	}

	// last temperature of each day (last 7 days) in Tel Aviv
	var tempVaryByPeriod []struct {
		Period      string  `json:"period"`
		Temperature float32 `json:"temperature"`
	}
	client.Sensor.Query().
		Where(sensor.SensorID(s2.id)).
		Order(ent.Desc("1")).
		GroupBy("1").
		Aggregate(func(selector *entsql.Selector) string {
			return "time_bucket('7 days', " + sensor.FieldCreateTime + ") AS period"
		}, func(selector *entsql.Selector) string {
			return "last(" + sensor.FieldValue + ", " + sensor.FieldCreateTime + ") AS temperature"
		}).ScanX(ctx, &tempVaryByPeriod)
	for _, r := range tempVaryByPeriod {
		_ = r
		// fmt.Println(r.Period, r.Temperature)
	}

	//
}

type tempSensor struct {
	id   string
	long float64
	lat  float64
}

func (s tempSensor) temperature() float64 {
	return -10 + rand.Float64()*(40-0)
}

func newTempSensor(long, lat float64) tempSensor {
	return tempSensor{
		id:   xid.New().String(),
		long: long,
		lat:  lat,
	}
}
func recordTemperature(entClient *ent.Client, n int, s tempSensor) {
	ctx := context.Background()
	t := time.Now()
	for i := 0; i < n; i++ {
		entClient.Sensor.Create().
			SetSensorID(s.id).
			SetType("temperature").
			SetLongitude(s.long).
			SetLatitude(s.lat).
			SetValue(s.temperature()).
			SetCreateTime(t).
			SaveX(ctx)
		t = t.Add(-(time.Duration(i) * time.Minute))
	}
}

func WithoutPrimaryKey(tableName string) schema.MigrateOption {
	return schema.WithHooks(func(next schema.Creator) schema.Creator {
		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
			for i := range tables {
				if tables[i].Name == tableName {
					tables[i].PrimaryKey = nil
				}
			}
			return next.Create(ctx, tables...)
		})
	})
}

func EnableTimescaleDBOption(db *sql.DB) schema.MigrateOption {
	return schema.WithHooks(func(next schema.Creator) schema.Creator {
		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
			_, err := db.ExecContext(ctx, `CREATE EXTENSION IF NOT EXISTS timescaledb WITH SCHEMA public;`)
			if err != nil {
				return fmt.Errorf("enable timescaledb extension: %w", err)
			}
			return next.Create(ctx, tables...)
		})
	})
}

func CreateHypertable(db *sql.DB, relation string, timeColumnName string) schema.MigrateOption {
	return schema.WithHooks(func(next schema.Creator) schema.Creator {
		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
			err := next.Create(ctx, tables...)
			if err != nil {
				return fmt.Errorf("executing migrations: %w", err)
			}
			_, err = db.ExecContext(ctx, `SELECT create_hypertable($1, $2, if_not_exists => TRUE);`, relation, timeColumnName)
			if err != nil {
				return fmt.Errorf("create_hypertable %q: %w", relation, err)
			}
			return nil
		})
	})
}
