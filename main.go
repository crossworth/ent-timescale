package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"ent-timescale/ent"
	"ent-timescale/ent/migrate"
	"ent-timescale/ent/sensor"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

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
		EnableTimescaleDBOption(db),
		WithoutPrimaryKey(sensor.Table),
		CreateHypertable(db, sensor.Table, sensor.FieldCreateTime),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// delete the data present on the database on each run
	client.Sensor.Delete().ExecX(ctx)

	// create a few records
	s1 := New("sensor_1")
	s2 := New("sensor_2")
	s3 := New("sensor_3")
	s4 := New("sensor_4")
	registerRecords(client, 10_000, s1)
	registerRecords(client, 10_000, s2)
	registerRecords(client, 10_000, s3)
	registerRecords(client, 10_000, s4)

	var possibleMachines []struct {
		Period string
		Avg    float64
		Last   float64
	}

	// without feature flags
	withAlias := func(alias string, aggregate func(selector *entsql.Selector) string) func(selector *entsql.Selector) string {
		return func(selector *entsql.Selector) string {
			return aggregate(selector) + " AS " + alias
		}
	}
	avg := func(selector *entsql.Selector) string {
		return "ABS(AVG(" + sensor.FieldTemperature + ") - AVG(" + sensor.FieldElectricCurrent + "))"
	}
	last := func(selector *entsql.Selector) string {
		return "ABS(LAST(" + sensor.FieldTemperature + "," + sensor.FieldCreateTime + ") - LAST(" + sensor.FieldElectricCurrent + "," + sensor.FieldCreateTime + "))"
	}
	client.Debug().Sensor.Query().
		Limit(10).
		Order(ent.Desc("1")).
		Where(func(selector *entsql.Selector) {
			selector.Having(
				entsql.Or(
					entsql.GT(avg(selector), 40),
					entsql.GT(last(selector), 40),
				),
			)
		}).
		GroupBy("1").
		Aggregate(func(selector *entsql.Selector) string {
			return "time_bucket('3 minutes', " + sensor.FieldCreateTime + ") AS period"
		}, withAlias("avg", avg), withAlias("last", last)).
		ScanX(ctx, &possibleMachines)
	for _, r := range possibleMachines {
		fmt.Println(r.Period, r.Avg, r.Last)
	}

	possibleMachines = nil

	// using feature flag modify
	client.Debug().Sensor.Query().Modify(func(s *entsql.Selector) {
		s.SelectExpr(entsql.Raw("time_bucket('3 minutes', create_time) AS period, ABS(AVG(temperature) - AVG(electric_current)) AS avg, ABS(LAST(temperature, create_time) - LAST(electric_current, create_time)) AS last"))
		s.GroupBy("period")
		s.Having(entsql.P(func(builder *entsql.Builder) {
			builder.WriteString("ABS(AVG(temperature) - AVG(electric_current)) > 40 OR ABS(LAST(temperature, create_time) - LAST(electric_current, create_time)) > 40")
		}))
		s.OrderExpr(entsql.Raw("period DESC"))
		s.Limit(10)
	}).ScanX(ctx, &possibleMachines)
	for _, r := range possibleMachines {
		fmt.Println(r.Period, r.Avg, r.Last)
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
