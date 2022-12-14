// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-timescale/ent/sensor"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SensorCreate is the builder for creating a Sensor entity.
type SensorCreate struct {
	config
	mutation *SensorMutation
	hooks    []Hook
}

// SetSensorID sets the "sensor_id" field.
func (sc *SensorCreate) SetSensorID(s string) *SensorCreate {
	sc.mutation.SetSensorID(s)
	return sc
}

// SetTemperature sets the "temperature" field.
func (sc *SensorCreate) SetTemperature(f float64) *SensorCreate {
	sc.mutation.SetTemperature(f)
	return sc
}

// SetElectricCurrent sets the "electric_current" field.
func (sc *SensorCreate) SetElectricCurrent(f float64) *SensorCreate {
	sc.mutation.SetElectricCurrent(f)
	return sc
}

// SetCreateTime sets the "create_time" field.
func (sc *SensorCreate) SetCreateTime(t time.Time) *SensorCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetID sets the "id" field.
func (sc *SensorCreate) SetID(u uuid.UUID) *SensorCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SensorCreate) SetNillableID(u *uuid.UUID) *SensorCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the SensorMutation object of the builder.
func (sc *SensorCreate) Mutation() *SensorMutation {
	return sc.mutation
}

// Save creates the Sensor in the database.
func (sc *SensorCreate) Save(ctx context.Context) (*Sensor, error) {
	var (
		err  error
		node *Sensor
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Sensor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SensorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SensorCreate) SaveX(ctx context.Context) *Sensor {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SensorCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SensorCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SensorCreate) defaults() {
	if _, ok := sc.mutation.ID(); !ok {
		v := sensor.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SensorCreate) check() error {
	if _, ok := sc.mutation.SensorID(); !ok {
		return &ValidationError{Name: "sensor_id", err: errors.New(`ent: missing required field "Sensor.sensor_id"`)}
	}
	if _, ok := sc.mutation.Temperature(); !ok {
		return &ValidationError{Name: "temperature", err: errors.New(`ent: missing required field "Sensor.temperature"`)}
	}
	if _, ok := sc.mutation.ElectricCurrent(); !ok {
		return &ValidationError{Name: "electric_current", err: errors.New(`ent: missing required field "Sensor.electric_current"`)}
	}
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Sensor.create_time"`)}
	}
	return nil
}

func (sc *SensorCreate) sqlSave(ctx context.Context) (*Sensor, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (sc *SensorCreate) createSpec() (*Sensor, *sqlgraph.CreateSpec) {
	var (
		_node = &Sensor{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sensor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: sensor.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.SensorID(); ok {
		_spec.SetField(sensor.FieldSensorID, field.TypeString, value)
		_node.SensorID = value
	}
	if value, ok := sc.mutation.Temperature(); ok {
		_spec.SetField(sensor.FieldTemperature, field.TypeFloat64, value)
		_node.Temperature = value
	}
	if value, ok := sc.mutation.ElectricCurrent(); ok {
		_spec.SetField(sensor.FieldElectricCurrent, field.TypeFloat64, value)
		_node.ElectricCurrent = value
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(sensor.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	return _node, _spec
}

// SensorCreateBulk is the builder for creating many Sensor entities in bulk.
type SensorCreateBulk struct {
	config
	builders []*SensorCreate
}

// Save creates the Sensor entities in the database.
func (scb *SensorCreateBulk) Save(ctx context.Context) ([]*Sensor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sensor, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SensorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SensorCreateBulk) SaveX(ctx context.Context) []*Sensor {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SensorCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SensorCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
