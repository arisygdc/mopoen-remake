-- name: CreateMonitoringTerdaftar :exec
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, nama, keterangan) VALUES ($1, $2, $3, $4, $5);

-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES ($1, $2);

-- name: GetMonitoringTerdaftarByLokasi :many
SELECT * FROM monitoring_terdaftar WHERE lokasi_id = $1;

-- name: GetMonitoringData :many
SELECT value, dibuat_pada FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: GetMonitoringTerdaftar :one
SELECT * FROM monitoring_terdaftar WHERE id = $1;

-- name: GetMonTerdaftarFilterLokAndSensor :many
SELECT * FROM monitoring_terdaftar WHERE tipe_sensor_id = $1 AND lokasi_id = $2;

-- name: CountDataMonitoring :one
SELECT COUNT(1) AS all, 
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00:00.1' AND '15:00:00') AS morning,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '12:00:00.1' AND '15:00:00') AS noon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00:00.1' AND '18:00:00') AS afternoon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00:00.1' AND '24:00:00') AS night,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '00:00:00.1' AND '06:00:00') AS midnight
FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: AverageDataMonitoring :one
SELECT AVG(value)::FLOAT AS all,
AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00:00.1' AND '15:00:00')::FLOAT AS morning,
AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '12:00:00.1' AND '15:00:00')::FLOAT AS noon,
AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00:00.1' AND '18:00:00')::FLOAT AS afternoon,
AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00:00.1' AND '24:00:00')::FLOAT AS night,
AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '00:00:00.1' AND '06:00:00')::FLOAT AS midnight
FROM monitoring_data WHERE monitoring_terdaftar = $1;