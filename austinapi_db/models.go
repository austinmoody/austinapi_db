// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package austinapi_db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Heartrate struct {
	ID               int32
	Date             pgtype.Date
	High             pgtype.Int4
	Low              pgtype.Int4
	Average          pgtype.Int4
	CreatedTimestamp pgtype.Timestamp
	UpdatedTimestamp pgtype.Timestamp
}

type Preparedness struct {
	ID               int32
	Date             pgtype.Date
	Rating           pgtype.Int4
	CreatedTimestamp pgtype.Timestamp
	UpdatedTimestamp pgtype.Timestamp
}

type Sleep struct {
	ID               int32
	Date             pgtype.Date
	Rating           pgtype.Int4
	CreatedTimestamp pgtype.Timestamp
	UpdatedTimestamp pgtype.Timestamp
	TotalDuration    pgtype.Int8
}

type Spo2 struct {
	ID               int32
	Date             pgtype.Date
	AverageSpo2      pgtype.Float8
	CreatedTimestamp pgtype.Timestamp
	UpdatedTimestamp pgtype.Timestamp
}

type Stress struct {
	ID                 int32
	Date               pgtype.Date
	HighStressDuration pgtype.Int4
	CreatedTimestamp   pgtype.Timestamp
	UpdatedTimestamp   pgtype.Timestamp
}
