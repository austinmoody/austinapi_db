-- name: SaveSleep :exec
INSERT INTO sleep (date, rating, total_sleep, deep_sleep, light_sleep, rem_sleep) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (date) DO UPDATE SET total_sleep = EXCLUDED.total_sleep, rating = EXCLUDED.rating, light_sleep = EXCLUDED.light_sleep, deep_sleep = EXCLUDED.deep_sleep, rem_sleep = EXCLUDED.rem_sleep;

-- name: GetSleep :many
SELECT * FROM sleep WHERE id = $1;

-- name: GetSleepByDate :many
SELECT * FROM sleep WHERE date = $1;

-- name: GetSleepDateById :many
SELECT date FROM sleep WHERE id = $1;

-- name: ListSleep :many
SELECT * FROM (
    SELECT *,
           CAST(LAG(id) OVER (ORDER BY date DESC) AS UUID)  AS previous_id,
           CAST(LEAD(id) OVER (ORDER BY date DESC) AS UUID) AS next_id
    FROM sleep
) sleeps
ORDER BY date DESC
LIMIT $1
;

-- name: ListSleepNext :many
SELECT * FROM (
      SELECT *,
             CAST(LAG(id) OVER (ORDER BY date DESC) AS UUID)  AS previous_id,
             CAST(LEAD(id) OVER (ORDER BY date DESC) AS UUID) AS next_id
      FROM sleep
              ) sleeps
WHERE date <= (
    SELECT date FROM sleep AS SLP WHERE SLP.id = $1
)
ORDER BY date DESC
LIMIT $2
;

-- name: ListSleepPrevious :many
SELECT * FROM (
      SELECT *,
             CAST(LAG(id) OVER (ORDER BY date DESC) AS UUID)  AS previous_id,
             CAST(LEAD(id) OVER (ORDER BY date DESC) AS UUID) AS next_id
      FROM sleep
              ) sleeps
WHERE date >= (
    SELECT date FROM sleep AS SLP WHERE SLP.id = $1
)
ORDER BY date DESC
LIMIT $2
;
-- name: SavePreparedness :exec
INSERT INTO preparedness (date, rating) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET rating = EXCLUDED.rating;

-- name: SaveSpo2 :exec
INSERT INTO spo2 (date, average_spo2) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET average_spo2 = EXCLUDED.average_spo2;

-- name: SaveStress :exec
INSERT INTO stress (date, high_stress_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET high_stress_duration = EXCLUDED.high_stress_duration;

-- name: SaveHeartRate :exec
INSERT INTO heartrate (date, low, high, average) VALUES ($1, $2, $3, $4) ON CONFLICT (date) DO UPDATE SET low = EXCLUDED.low, high = EXCLUDED.high, average = EXCLUDED.average;