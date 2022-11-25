package main

import (
	"context"
	"math/rand"
	"time"

	"ent-timescale/ent"

	"github.com/google/uuid"
)

type fakeSensor struct {
	id string
}

func New(id string) fakeSensor {
	return fakeSensor{id: id}
}

func (f fakeSensor) Data() (id uuid.UUID, sensorID string, temperature float64, electricCurrent float64) {
	return uuid.New(), f.id, 30 + rand.Float64()*(70-0), 10 + rand.Float64()*(30-0)
}

func registerRecords(entClient *ent.Client, n int, s fakeSensor) {
	ctx := context.Background()
	t := time.Now()
	for i := 0; i < n; i++ {
		id, sensorID, temp, curr := s.Data()
		entClient.Sensor.Create().
			SetID(id).
			SetSensorID(sensorID).
			SetTemperature(temp).
			SetElectricCurrent(curr).
			SetCreateTime(t).
			SaveX(ctx)
		t = t.Add(1 * time.Minute)
	}
}
