-- name: CreateMonitoringTerdaftar :exec
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, nama, keterangan) VALUES ($1, $2, $3, $4, $5);

-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES ($1, $2);

-- name: GetMonitoringTerdaftarByLokasi :many
SELECT * FROM monitoring_terdaftar WHERE lokasi_id = $1;

-- name: GetMonitoringData :many
SELECT (value) FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: GetMonitoringTerdaftar :one
SELECT * FROM monitoring_terdaftar WHERE id = $1;