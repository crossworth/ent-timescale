// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent-timescale/ent/sensor"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Sensor is the model entity for the Sensor schema.
type Sensor struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SensorID holds the value of the "sensor_id" field.
	SensorID string `json:"sensor_id,omitempty"`
	// Temperature holds the value of the "temperature" field.
	Temperature float64 `json:"temperature,omitempty"`
	// ElectricCurrent holds the value of the "electric_current" field.
	ElectricCurrent float64 `json:"electric_current,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sensor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sensor.FieldTemperature, sensor.FieldElectricCurrent:
			values[i] = new(sql.NullFloat64)
		case sensor.FieldSensorID:
			values[i] = new(sql.NullString)
		case sensor.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case sensor.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sensor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sensor fields.
func (s *Sensor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sensor.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case sensor.FieldSensorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sensor_id", values[i])
			} else if value.Valid {
				s.SensorID = value.String
			}
		case sensor.FieldTemperature:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temperature", values[i])
			} else if value.Valid {
				s.Temperature = value.Float64
			}
		case sensor.FieldElectricCurrent:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field electric_current", values[i])
			} else if value.Valid {
				s.ElectricCurrent = value.Float64
			}
		case sensor.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Sensor.
// Note that you need to call Sensor.Unwrap() before calling this method if this Sensor
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sensor) Update() *SensorUpdateOne {
	return (&SensorClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Sensor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sensor) Unwrap() *Sensor {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sensor is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sensor) String() string {
	var builder strings.Builder
	builder.WriteString("Sensor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("sensor_id=")
	builder.WriteString(s.SensorID)
	builder.WriteString(", ")
	builder.WriteString("temperature=")
	builder.WriteString(fmt.Sprintf("%v", s.Temperature))
	builder.WriteString(", ")
	builder.WriteString("electric_current=")
	builder.WriteString(fmt.Sprintf("%v", s.ElectricCurrent))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sensors is a parsable slice of Sensor.
type Sensors []*Sensor

func (s Sensors) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
