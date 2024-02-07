// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package austinapi_db

import (
	"time"
)

type Heartrate struct {
	ID               int32
	Date             time.Time
	High             int
	Low              int
	Average          int
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}

type Preparedness struct {
	ID               int32
	Date             time.Time
	Rating           int
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}

type Sleep struct {
	ID               int32
	Date             time.Time
	Rating           int64
	TotalSleep       int
	DeepSleep        int
	LightSleep       int
	RemSleep         int
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}

type Spo2 struct {
	ID               int32
	Date             time.Time
	AverageSpo2      float64
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}

type Stress struct {
	ID                 int32
	Date               time.Time
	HighStressDuration int
	CreatedTimestamp   time.Time
	UpdatedTimestamp   time.Time
}
