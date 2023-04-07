-- name: CreateMonitoringTerdaftar :exec
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, nama, keterangan) VALUES ($1, $2, $3, $4, $5);

-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES ($1, $2);

-- name: GetMonitoringTerdaftarByLokasi :many
SELECT * FROM monitoring_terdaftar WHERE lokasi_id = $1;

-- name: GetAllMonitoringTerdaftar :many
SELECT * FROM monitoring_terdaftar;

-- name: GetMonitoringData :many
SELECT value, dibuat_pada FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: GetMonitoringTerdaftar :one
SELECT * FROM monitoring_terdaftar WHERE id = $1;

-- name: GetMonTerdaftarFilterLokAndSensor :many
SELECT * FROM monitoring_terdaftar WHERE tipe_sensor_id = $1 AND lokasi_id = $2;

-- name: CountDataMonitoring :one
SELECT COUNT(1) AS all, 
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00:00' AND '11:59') AS morning,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '12:00:00' AND '14:59') AS noon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00:00' AND '17:59') AS afternoon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00:00' AND '23:59') AS night,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '00:00:00' AND '05:59') AS midnight
FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: AverageDataMonitoring :one
SELECT COALESCE(AVG(value), 0)::FLOAT AS all,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00:00' AND '11:59'), 0)::FLOAT AS morning,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '12:00:00' AND '14:59'), 0)::FLOAT AS noon,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00:00' AND '17:59'), 0)::FLOAT AS afternoon,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00:00' AND '23:59'), 0)::FLOAT AS night,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '00:00:00' AND '05:59'), 0)::FLOAT AS midnight
FROM monitoring_data WHERE monitoring_terdaftar = $1;