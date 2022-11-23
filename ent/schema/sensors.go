package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Sensor holds the schema definition for the Sensor entity.
type Sensor struct {
	ent.Schema
}

// Fields of the Sensor.
func (Sensor) Fields() []ent.Field {
	incrementalEnabled := false

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().
			Annotations(entsql.Annotation{
				Incremental: &incrementalEnabled,
			}).Default(func() uuid.UUID {
			return uuid.New()
		}),
		field.String("sensor_id"),
		field.Float("longitude"),
		field.Float("latitude"),
		field.String("type"),
		field.Float("value"),
		field.Time("create_time"),
	}
}

// Indexes of the Sensor.
func (Sensor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sensor_id"),
		index.Fields("type"),
	}
}
