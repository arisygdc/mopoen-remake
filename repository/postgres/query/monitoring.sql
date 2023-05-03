-- name: CreateMonitoringTerdaftar :one
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, email, secret, author, nama, keterangan) 
VALUES (@id, @tipe_sensor_id, @lokasi_id, @email, @secret, @author, @nama, @keterangan) RETURNING *;

-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES 
(
    (
        SELECT monitoring_terdaftar.id 
        FROM monitoring_terdaftar 
        WHERE monitoring_terdaftar.id=$1 
        AND monitoring_terdaftar.secret=$3
    ), 
    $2
);

-- name: GetAllMonitoringTerdaftar :many
SELECT 
    mt.id as monitoring_id, 
    mt.tipe_sensor_id,
    concat(ts.tipe, ' (', ts.satuan, ')' )::text as tipe_sensor, 
    mt.nama, 
    mt.keterangan, 
    pv.nama AS provinsi, 
    kc.nama AS kecamatan, 
    kb.nama AS kabupaten, 
    d.nama AS desa
FROM 
    monitoring_terdaftar mt left join tipe_sensor ts on mt.tipe_sensor_id = ts.id
    left join desa d on d.id = mt.lokasi_id 
    left join kecamatan kc on d.kecamatan_id = kc.id 
    left join kabupaten kb on kc.kabupaten_id = kb.id 
    left join provinsi pv on kb.provinsi_id = pv.id;

-- name: GetMonitoringTerdaftarFilter :many
SELECT 
    mt.id as monitoring_id, 
    mt.tipe_sensor_id, 
    concat(ts.tipe, ' (', ts.satuan, ')' )::text as tipe_sensor, 
    mt.nama, 
    mt.keterangan, 
    pv.nama AS provinsi, 
    kc.nama AS kecamatan, 
    kb.nama AS kabupaten, 
    d.nama AS desa
FROM 
    monitoring_terdaftar mt left join tipe_sensor ts on mt.tipe_sensor_id = ts.id
    left join desa d on d.id = mt.lokasi_id 
    left join kecamatan kc on d.kecamatan_id = kc.id 
    left join kabupaten kb on kc.kabupaten_id = kb.id 
    left join provinsi pv on kb.provinsi_id = pv.id
WHERE
CASE 
    WHEN $1 > 0 AND $2 > 0 THEN
        mt.tipe_sensor_id = $1 AND mt.lokasi_id = $2
    WHEN $1 > 0 THEN
        mt.tipe_sensor_id = $1
    WHEN $2 > 0 THEN
        mt.lokasi_id = $2
END;

-- name: GetMonitoringData :many
SELECT value, dibuat_pada FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: GetMonitoringTerdaftar :one
SELECT 
    mt.id as monitoring_id, 
    mt.tipe_sensor_id, 
    concat(ts.tipe, ' (', ts.satuan, ')' )::text as tipe_sensor, 
    mt.nama, 
    mt.keterangan, 
    concat(d.nama, ', ', kc.nama, ', ', kb.nama, ', ', pv.nama)::varchar as address
FROM 
    monitoring_terdaftar mt left join tipe_sensor ts on mt.tipe_sensor_id = ts.id
    left join desa d on d.id = mt.lokasi_id 
    left join kecamatan kc on d.kecamatan_id = kc.id 
    left join kabupaten kb on kc.kabupaten_id = kb.id 
    left join provinsi pv on kb.provinsi_id = pv.id 
WHERE
    mt.id = $1;

-- name: GetMonTerdaftarFilterLokAndSensor :many
SELECT id, tipe_sensor_id, lokasi_id, email, author, nama, keterangan FROM monitoring_terdaftar WHERE tipe_sensor_id = $1 AND lokasi_id = $2;

-- name: CountDataMonitoring :one
SELECT 
    COUNT(1) AS all, 
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 6 AND 10) AS morning,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 11 AND 14) AS noon,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 15 AND 17) AS afternoon,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) >= 18 OR EXTRACT(HOUR FROM dibuat_pada) < 6) AS night
FROM 
    monitoring_data
WHERE 
    monitoring_terdaftar = $1;

-- name: AverageDataMonitoring :one
SELECT 
    COALESCE(AVG(value), 0)::FLOAT AS all,
    COALESCE(AVG(value) 
        FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 6 AND 10), 0)::FLOAT AS morning,
    COALESCE(AVG(value) 
        FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 11 AND 14), 0)::FLOAT AS noon,
    COALESCE(AVG(value) 
        FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 15 AND 17), 0)::FLOAT AS afternoon,
    COALESCE(AVG(value) 
        FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) >= 18 OR EXTRACT(HOUR FROM dibuat_pada) < 6), 0)::FLOAT AS night
FROM 
    monitoring_data 
WHERE 
    monitoring_terdaftar = $1;