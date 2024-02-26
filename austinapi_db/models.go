// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package austinapi_db

import (
	"time"
)

type Heartrate struct {
	ID               int64     `json:"id"`
	Date             time.Time `json:"date"`
	High             int       `json:"high"`
	Low              int       `json:"low"`
	Average          int       `json:"average"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type Readyscore struct {
	ID               int64     `json:"id"`
	Date             time.Time `json:"date"`
	Score            int       `json:"score"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type Sleep struct {
	ID               int64     `json:"id"`
	Date             time.Time `json:"date"`
	Rating           int64     `json:"rating"`
	TotalSleep       int       `json:"total_sleep"`
	DeepSleep        int       `json:"deep_sleep"`
	LightSleep       int       `json:"light_sleep"`
	RemSleep         int       `json:"rem_sleep"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type Spo2 struct {
	ID               int64     `json:"id"`
	Date             time.Time `json:"date"`
	AverageSpo2      float64   `json:"average_spo2"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type Stress struct {
	ID                 int64     `json:"id"`
	Date               time.Time `json:"date"`
	HighStressDuration int64     `json:"high_stress_duration"`
	CreatedTimestamp   time.Time `json:"created_timestamp"`
	UpdatedTimestamp   time.Time `json:"updated_timestamp"`
}
