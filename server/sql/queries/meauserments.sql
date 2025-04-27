-- name: CreateMeasurement :one
INSERT INTO measurements (id, measurement_time, humidity, temperature)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMeasurements :many
SELECT * FROM measurements;
