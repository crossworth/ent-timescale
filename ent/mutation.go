// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-timescale/ent/predicate"
	"ent-timescale/ent/sensor"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSensor = "Sensor"
)

// SensorMutation represents an operation that mutates the Sensor nodes in the graph.
type SensorMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	sensor_id     *string
	longitude     *float64
	addlongitude  *float64
	latitude      *float64
	addlatitude   *float64
	_type         *string
	value         *float64
	addvalue      *float64
	create_time   *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Sensor, error)
	predicates    []predicate.Sensor
}

var _ ent.Mutation = (*SensorMutation)(nil)

// sensorOption allows management of the mutation configuration using functional options.
type sensorOption func(*SensorMutation)

// newSensorMutation creates new mutation for the Sensor entity.
func newSensorMutation(c config, op Op, opts ...sensorOption) *SensorMutation {
	m := &SensorMutation{
		config:        c,
		op:            op,
		typ:           TypeSensor,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSensorID sets the ID field of the mutation.
func withSensorID(id uuid.UUID) sensorOption {
	return func(m *SensorMutation) {
		var (
			err   error
			once  sync.Once
			value *Sensor
		)
		m.oldValue = func(ctx context.Context) (*Sensor, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Sensor.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSensor sets the old Sensor of the mutation.
func withSensor(node *Sensor) sensorOption {
	return func(m *SensorMutation) {
		m.oldValue = func(context.Context) (*Sensor, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SensorMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SensorMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Sensor entities.
func (m *SensorMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SensorMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SensorMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Sensor.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetSensorID sets the "sensor_id" field.
func (m *SensorMutation) SetSensorID(s string) {
	m.sensor_id = &s
}

// SensorID returns the value of the "sensor_id" field in the mutation.
func (m *SensorMutation) SensorID() (r string, exists bool) {
	v := m.sensor_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSensorID returns the old "sensor_id" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldSensorID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSensorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSensorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSensorID: %w", err)
	}
	return oldValue.SensorID, nil
}

// ResetSensorID resets all changes to the "sensor_id" field.
func (m *SensorMutation) ResetSensorID() {
	m.sensor_id = nil
}

// SetLongitude sets the "longitude" field.
func (m *SensorMutation) SetLongitude(f float64) {
	m.longitude = &f
	m.addlongitude = nil
}

// Longitude returns the value of the "longitude" field in the mutation.
func (m *SensorMutation) Longitude() (r float64, exists bool) {
	v := m.longitude
	if v == nil {
		return
	}
	return *v, true
}

// OldLongitude returns the old "longitude" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldLongitude(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLongitude is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLongitude requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLongitude: %w", err)
	}
	return oldValue.Longitude, nil
}

// AddLongitude adds f to the "longitude" field.
func (m *SensorMutation) AddLongitude(f float64) {
	if m.addlongitude != nil {
		*m.addlongitude += f
	} else {
		m.addlongitude = &f
	}
}

// AddedLongitude returns the value that was added to the "longitude" field in this mutation.
func (m *SensorMutation) AddedLongitude() (r float64, exists bool) {
	v := m.addlongitude
	if v == nil {
		return
	}
	return *v, true
}

// ResetLongitude resets all changes to the "longitude" field.
func (m *SensorMutation) ResetLongitude() {
	m.longitude = nil
	m.addlongitude = nil
}

// SetLatitude sets the "latitude" field.
func (m *SensorMutation) SetLatitude(f float64) {
	m.latitude = &f
	m.addlatitude = nil
}

// Latitude returns the value of the "latitude" field in the mutation.
func (m *SensorMutation) Latitude() (r float64, exists bool) {
	v := m.latitude
	if v == nil {
		return
	}
	return *v, true
}

// OldLatitude returns the old "latitude" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldLatitude(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLatitude is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLatitude requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLatitude: %w", err)
	}
	return oldValue.Latitude, nil
}

// AddLatitude adds f to the "latitude" field.
func (m *SensorMutation) AddLatitude(f float64) {
	if m.addlatitude != nil {
		*m.addlatitude += f
	} else {
		m.addlatitude = &f
	}
}

// AddedLatitude returns the value that was added to the "latitude" field in this mutation.
func (m *SensorMutation) AddedLatitude() (r float64, exists bool) {
	v := m.addlatitude
	if v == nil {
		return
	}
	return *v, true
}

// ResetLatitude resets all changes to the "latitude" field.
func (m *SensorMutation) ResetLatitude() {
	m.latitude = nil
	m.addlatitude = nil
}

// SetType sets the "type" field.
func (m *SensorMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *SensorMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *SensorMutation) ResetType() {
	m._type = nil
}

// SetValue sets the "value" field.
func (m *SensorMutation) SetValue(f float64) {
	m.value = &f
	m.addvalue = nil
}

// Value returns the value of the "value" field in the mutation.
func (m *SensorMutation) Value() (r float64, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old "value" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldValue(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldValue is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// AddValue adds f to the "value" field.
func (m *SensorMutation) AddValue(f float64) {
	if m.addvalue != nil {
		*m.addvalue += f
	} else {
		m.addvalue = &f
	}
}

// AddedValue returns the value that was added to the "value" field in this mutation.
func (m *SensorMutation) AddedValue() (r float64, exists bool) {
	v := m.addvalue
	if v == nil {
		return
	}
	return *v, true
}

// ResetValue resets all changes to the "value" field.
func (m *SensorMutation) ResetValue() {
	m.value = nil
	m.addvalue = nil
}

// SetCreateTime sets the "create_time" field.
func (m *SensorMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *SensorMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *SensorMutation) ResetCreateTime() {
	m.create_time = nil
}

// Where appends a list predicates to the SensorMutation builder.
func (m *SensorMutation) Where(ps ...predicate.Sensor) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SensorMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Sensor).
func (m *SensorMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SensorMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.sensor_id != nil {
		fields = append(fields, sensor.FieldSensorID)
	}
	if m.longitude != nil {
		fields = append(fields, sensor.FieldLongitude)
	}
	if m.latitude != nil {
		fields = append(fields, sensor.FieldLatitude)
	}
	if m._type != nil {
		fields = append(fields, sensor.FieldType)
	}
	if m.value != nil {
		fields = append(fields, sensor.FieldValue)
	}
	if m.create_time != nil {
		fields = append(fields, sensor.FieldCreateTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SensorMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sensor.FieldSensorID:
		return m.SensorID()
	case sensor.FieldLongitude:
		return m.Longitude()
	case sensor.FieldLatitude:
		return m.Latitude()
	case sensor.FieldType:
		return m.GetType()
	case sensor.FieldValue:
		return m.Value()
	case sensor.FieldCreateTime:
		return m.CreateTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SensorMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sensor.FieldSensorID:
		return m.OldSensorID(ctx)
	case sensor.FieldLongitude:
		return m.OldLongitude(ctx)
	case sensor.FieldLatitude:
		return m.OldLatitude(ctx)
	case sensor.FieldType:
		return m.OldType(ctx)
	case sensor.FieldValue:
		return m.OldValue(ctx)
	case sensor.FieldCreateTime:
		return m.OldCreateTime(ctx)
	}
	return nil, fmt.Errorf("unknown Sensor field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sensor.FieldSensorID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSensorID(v)
		return nil
	case sensor.FieldLongitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLongitude(v)
		return nil
	case sensor.FieldLatitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLatitude(v)
		return nil
	case sensor.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case sensor.FieldValue:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	case sensor.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	}
	return fmt.Errorf("unknown Sensor field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SensorMutation) AddedFields() []string {
	var fields []string
	if m.addlongitude != nil {
		fields = append(fields, sensor.FieldLongitude)
	}
	if m.addlatitude != nil {
		fields = append(fields, sensor.FieldLatitude)
	}
	if m.addvalue != nil {
		fields = append(fields, sensor.FieldValue)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SensorMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case sensor.FieldLongitude:
		return m.AddedLongitude()
	case sensor.FieldLatitude:
		return m.AddedLatitude()
	case sensor.FieldValue:
		return m.AddedValue()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorMutation) AddField(name string, value ent.Value) error {
	switch name {
	case sensor.FieldLongitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLongitude(v)
		return nil
	case sensor.FieldLatitude:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLatitude(v)
		return nil
	case sensor.FieldValue:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValue(v)
		return nil
	}
	return fmt.Errorf("unknown Sensor numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SensorMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SensorMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SensorMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Sensor nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SensorMutation) ResetField(name string) error {
	switch name {
	case sensor.FieldSensorID:
		m.ResetSensorID()
		return nil
	case sensor.FieldLongitude:
		m.ResetLongitude()
		return nil
	case sensor.FieldLatitude:
		m.ResetLatitude()
		return nil
	case sensor.FieldType:
		m.ResetType()
		return nil
	case sensor.FieldValue:
		m.ResetValue()
		return nil
	case sensor.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	}
	return fmt.Errorf("unknown Sensor field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SensorMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SensorMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SensorMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SensorMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SensorMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SensorMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SensorMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Sensor unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SensorMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Sensor edge %s", name)
}