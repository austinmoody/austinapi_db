-- name: SaveSleep :exec
INSERT INTO sleep (date, rating, total_duration) VALUES ($1, $2, $3) ON CONFLICT (date) DO UPDATE SET total_duration = EXCLUDED.total_duration, rating = EXCLUDED.rating;
-- TODO - Rethink naming here.  Add indicates we'd be potentially adding to what is there? maybe thereis some way w/ insert to add?  So if we insert a duration we'd add to what is there and increment

-- name: SaveSleepRating :exec
INSERT INTO sleep (date, rating) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET rating = EXCLUDED.rating;

-- name: SaveSleepDuration :exec
INSERT INTO sleep (date, total_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET total_duration = EXCLUDED.total_duration;

-- name: GetSleep :one
SELECT * FROM sleep WHERE id = $1;

-- name: GetSleepByDate :one
SELECT * FROM sleep WHERE date = $1;

-- name: SavePreparedness :exec
INSERT INTO preparedness (date, rating) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET rating = EXCLUDED.rating;

-- name: SaveSpo2 :exec
INSERT INTO spo2 (date, average_spo2) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET average_spo2 = EXCLUDED.average_spo2;

-- name: SaveStress :exec
INSERT INTO stress (date, high_stress_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET high_stress_duration = EXCLUDED.high_stress_duration;

-- name: SaveHeartRate :exec
INSERT INTO heartrate (date, low, high, average) VALUES ($1, $2, $3, $4) ON CONFLICT (date) DO UPDATE SET low = EXCLUDED.low, high = EXCLUDED.high, average = EXCLUDED.average;