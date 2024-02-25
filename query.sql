-- name: SaveSleep :exec
INSERT INTO sleep (date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (date) DO UPDATE SET total_sleep = EXCLUDED.total_sleep, rating = EXCLUDED.rating, light_sleep = EXCLUDED.light_sleep, deep_sleep = EXCLUDED.deep_sleep, rem_sleep = EXCLUDED.rem_sleep;

-- name: GetSleep :many
SELECT *
FROM sleep
WHERE id = $1;

-- name: GetSleepByDate :many
SELECT *
FROM sleep
WHERE date = $1;

-- name: GetSleeps :many
SELECT *
FROM sleep
ORDER BY date DESC
LIMIT sqlc.arg(row_limit) OFFSET sqlc.arg(row_offset);

-- name: SaveReadyScore :exec
INSERT INTO readyscore (date, score) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET score = EXCLUDED.score;

-- name: GetReadyScore :many
SELECT *
FROM readyscore
WHERE id = $1;

-- name: GetReadyScoreByDate :many
SELECT *
FROM readyscore
WHERE date = $1;

-- name: GetReadyScores :many
SELECT *
FROM readyscore
ORDER BY date DESC
LIMIT sqlc.arg(row_limit) OFFSET sqlc.arg(row_offset);

-- name: SaveSpo2 :exec
INSERT INTO spo2 (date, average_spo2) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET average_spo2 = EXCLUDED.average_spo2;

-- name: GetSpo2 :many
SELECT *
FROM spo2
WHERE id = $1;

-- name: GetSpo2ByDate :many
SELECT *
FROM spo2
WHERE date = $1;

-- name: GetSpo2s :many
SELECT *
FROM spo2
ORDER BY date DESC
LIMIT sqlc.arg(row_limit) OFFSET sqlc.arg(row_offset);

-- name: SaveStress :exec
INSERT INTO stress (date, high_stress_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET high_stress_duration = EXCLUDED.high_stress_duration;

-- name: GetStress :many
SELECT *
FROM stress
WHERE id = $1;

-- name: GetStressByDate :many
SELECT *
FROM stress
WHERE date = $1;

-- name: GetStresses :many
SELECT *
FROM stress
ORDER BY date DESC
LIMIT sqlc.arg(row_limit) OFFSET sqlc.arg(row_offset);

-- name: SaveHeartRate :exec
INSERT INTO heartrate (date, low, high, average) VALUES ($1, $2, $3, $4) ON CONFLICT (date) DO UPDATE SET low = EXCLUDED.low, high = EXCLUDED.high, average = EXCLUDED.average;

-- name: GetHeartRate :many
SELECT *
FROM heartrate
WHERE id = $1;

-- name: GetHeartRateByDate :many
SELECT *
FROM heartrate
WHERE date = $1;

-- name: GetHeartRates :many
SELECT *
FROM heartrate
ORDER BY date DESC
LIMIT sqlc.arg(row_limit) OFFSET sqlc.arg(row_offset);
