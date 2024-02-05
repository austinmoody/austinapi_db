-- name: AddSleep :exec
INSERT INTO sleep (date, rating, total_duration) VALUES ($1, $2, $3) ON CONFLICT (date) DO UPDATE SET total_duration = EXCLUDED.total_duration, rating = EXCLUDED.rating;

-- name: AddSleepRating :exec
INSERT INTO sleep (date, rating) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET rating = EXCLUDED.rating;

-- name: AddSleepDuration :exec
INSERT INTO sleep (date, total_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET total_duration = EXCLUDED.total_duration;

-- name: AddPreparedness :exec
INSERT INTO preparedness (date, rating) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET rating = EXCLUDED.rating;

-- name: AddSpo2 :exec
INSERT INTO spo2 (date, average_spo2) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET average_spo2 = EXCLUDED.average_spo2;

-- name: AddStress :exec
INSERT INTO stress (date, high_stress_duration) VALUES ($1, $2) ON CONFLICT (date) DO UPDATE SET high_stress_duration = EXCLUDED.high_stress_duration;