// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package austinapi_db

import (
	"context"
	"time"
)

const getReadyScore = `-- name: GetReadyScore :many
SELECT id, date, score, created_timestamp, updated_timestamp FROM readyscore WHERE id = $1
`

func (q *Queries) GetReadyScore(ctx context.Context, id int64) ([]Readyscore, error) {
	rows, err := q.db.Query(ctx, getReadyScore, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Readyscore{}
	for rows.Next() {
		var i Readyscore
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Score,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadyScoreByDate = `-- name: GetReadyScoreByDate :many
SELECT id, date, score, created_timestamp, updated_timestamp FROM readyscore WHERE date = $1
`

func (q *Queries) GetReadyScoreByDate(ctx context.Context, date time.Time) ([]Readyscore, error) {
	rows, err := q.db.Query(ctx, getReadyScoreByDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Readyscore{}
	for rows.Next() {
		var i Readyscore
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Score,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadyScoreById = `-- name: GetReadyScoreById :many
SELECT id, date, score, created_timestamp, updated_timestamp FROM readyscore WHERE id = $1
`

func (q *Queries) GetReadyScoreById(ctx context.Context, id int64) ([]Readyscore, error) {
	rows, err := q.db.Query(ctx, getReadyScoreById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Readyscore{}
	for rows.Next() {
		var i Readyscore
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Score,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadyScores = `-- name: GetReadyScores :many
SELECT id, date, score, created_timestamp, updated_timestamp, previous_id, next_id FROM (
      SELECT id, date, score, created_timestamp, updated_timestamp,
             CAST(COALESCE(LAG(id) OVER (ORDER BY date DESC), -1) AS BIGINT) AS previous_id,
             CAST(COALESCE(LEAD(id) OVER (ORDER BY date DESC), -1) AS BIGINT) AS next_id
      FROM readyscore
) listreadyscores
WHERE CASE
          WHEN 'NEXT' = $1::text THEN date <= (SELECT date FROM readyscore AS SLP WHERE SLP.id = $2)
          WHEN 'PREVIOUS' = $1::text THEN date >= (SELECT date FROM readyscore AS SLP WHERE SLP.id = $2)
          ELSE true
          END
ORDER BY date DESC
LIMIT $3
`

type GetReadyScoresParams struct {
	QueryType string
	InputID   int64
	RowLimit  int32
}

type GetReadyScoresRow struct {
	ID               int64
	Date             time.Time
	Score            int
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
	PreviousID       int64
	NextID           int64
}

func (q *Queries) GetReadyScores(ctx context.Context, arg GetReadyScoresParams) ([]GetReadyScoresRow, error) {
	rows, err := q.db.Query(ctx, getReadyScores, arg.QueryType, arg.InputID, arg.RowLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetReadyScoresRow{}
	for rows.Next() {
		var i GetReadyScoresRow
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Score,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
			&i.PreviousID,
			&i.NextID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSleep = `-- name: GetSleep :many
SELECT id, date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep, created_timestamp, updated_timestamp FROM sleep WHERE id = $1
`

func (q *Queries) GetSleep(ctx context.Context, id int64) ([]Sleep, error) {
	rows, err := q.db.Query(ctx, getSleep, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sleep{}
	for rows.Next() {
		var i Sleep
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Rating,
			&i.TotalSleep,
			&i.DeepSleep,
			&i.LightSleep,
			&i.RemSleep,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSleepByDate = `-- name: GetSleepByDate :many
SELECT id, date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep, created_timestamp, updated_timestamp FROM sleep WHERE date = $1
`

func (q *Queries) GetSleepByDate(ctx context.Context, date time.Time) ([]Sleep, error) {
	rows, err := q.db.Query(ctx, getSleepByDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Sleep{}
	for rows.Next() {
		var i Sleep
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Rating,
			&i.TotalSleep,
			&i.DeepSleep,
			&i.LightSleep,
			&i.RemSleep,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSleepDateById = `-- name: GetSleepDateById :many
SELECT date FROM sleep WHERE id = $1
`

func (q *Queries) GetSleepDateById(ctx context.Context, id int64) ([]time.Time, error) {
	rows, err := q.db.Query(ctx, getSleepDateById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []time.Time{}
	for rows.Next() {
		var date time.Time
		if err := rows.Scan(&date); err != nil {
			return nil, err
		}
		items = append(items, date)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveHeartRate = `-- name: SaveHeartRate :exec
INSERT INTO heartrate (date, low, high, average) VALUES ($1, $2, $3, $4) ON CONFLICT (date) DO UPDATE SET low = EXCLUDED.low, high = EXCLUDED.high, average = EXCLUDED.average
`

type SaveHeartRateParams struct {
	Date    time.Time
	Low     int
	High    int
	Average int
}

func (q *Queries) SaveHeartRate(ctx context.Context, arg SaveHeartRateParams) error {
	_, err := q.db.Exec(ctx, saveHeartRate,
		arg.Date,
		arg.Low,
		arg.High,
		arg.Average,
	)
	return err
}

const saveReadyScore = `-- name: SaveReadyScore :exec
INSERT INTO readyscore (date, score) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET score = EXCLUDED.score
`

type SaveReadyScoreParams struct {
	Date  time.Time
	Score int
}

func (q *Queries) SaveReadyScore(ctx context.Context, arg SaveReadyScoreParams) error {
	_, err := q.db.Exec(ctx, saveReadyScore, arg.Date, arg.Score)
	return err
}

const saveSleep = `-- name: SaveSleep :exec
INSERT INTO sleep (date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (date) DO UPDATE SET total_sleep = EXCLUDED.total_sleep, rating = EXCLUDED.rating, light_sleep = EXCLUDED.light_sleep, deep_sleep = EXCLUDED.deep_sleep, rem_sleep = EXCLUDED.rem_sleep
`

type SaveSleepParams struct {
	Date       time.Time
	Rating     int64
	TotalSleep int
	DeepSleep  int
	LightSleep int
	RemSleep   int
}

func (q *Queries) SaveSleep(ctx context.Context, arg SaveSleepParams) error {
	_, err := q.db.Exec(ctx, saveSleep,
		arg.Date,
		arg.Rating,
		arg.TotalSleep,
		arg.DeepSleep,
		arg.LightSleep,
		arg.RemSleep,
	)
	return err
}

const saveSpo2 = `-- name: SaveSpo2 :exec
INSERT INTO spo2 (date, average_spo2) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET average_spo2 = EXCLUDED.average_spo2
`

type SaveSpo2Params struct {
	Date        time.Time
	AverageSpo2 float64
}

func (q *Queries) SaveSpo2(ctx context.Context, arg SaveSpo2Params) error {
	_, err := q.db.Exec(ctx, saveSpo2, arg.Date, arg.AverageSpo2)
	return err
}

const saveStress = `-- name: SaveStress :exec
INSERT INTO stress (date, high_stress_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET high_stress_duration = EXCLUDED.high_stress_duration
`

type SaveStressParams struct {
	Date               time.Time
	HighStressDuration int
}

func (q *Queries) SaveStress(ctx context.Context, arg SaveStressParams) error {
	_, err := q.db.Exec(ctx, saveStress, arg.Date, arg.HighStressDuration)
	return err
}

const sleeps = `-- name: Sleeps :many
SELECT id, date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep, created_timestamp, updated_timestamp, previous_id, next_id FROM (
  SELECT id, date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep, created_timestamp, updated_timestamp,
         CAST(COALESCE(LAG(id) OVER (ORDER BY date DESC), -1) AS BIGINT) AS previous_id,
         CAST(COALESCE(LEAD(id) OVER (ORDER BY date DESC), -1) AS BIGINT) AS next_id
  FROM sleep
) sleeps
WHERE CASE
          WHEN 'NEXT' = $1::text THEN date <= (SELECT date FROM sleep AS SLP WHERE SLP.id = $2)
          WHEN 'PREVIOUS' = $1::text THEN date >= (SELECT date FROM sleep AS SLP WHERE SLP.id = $2)
          ELSE true
          END
ORDER BY date DESC
LIMIT $3
`

type SleepsParams struct {
	QueryType string
	InputID   int64
	RowLimit  int32
}

type SleepsRow struct {
	ID               int64
	Date             time.Time
	Rating           int64
	TotalSleep       int
	DeepSleep        int
	LightSleep       int
	RemSleep         int
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
	PreviousID       int64
	NextID           int64
}

func (q *Queries) Sleeps(ctx context.Context, arg SleepsParams) ([]SleepsRow, error) {
	rows, err := q.db.Query(ctx, sleeps, arg.QueryType, arg.InputID, arg.RowLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SleepsRow{}
	for rows.Next() {
		var i SleepsRow
		if err := rows.Scan(
			&i.ID,
			&i.Date,
			&i.Rating,
			&i.TotalSleep,
			&i.DeepSleep,
			&i.LightSleep,
			&i.RemSleep,
			&i.CreatedTimestamp,
			&i.UpdatedTimestamp,
			&i.PreviousID,
			&i.NextID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
