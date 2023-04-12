-- name: CreateMonitoringTerdaftar :one
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, email, author, nama, keterangan) 
VALUES (@id, @tipe_sensor_id, @lokasi_id, @email, @author, @nama, @keterangan) RETURNING *;

-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES ($1, $2);

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
SELECT * FROM monitoring_terdaftar WHERE tipe_sensor_id = $1 AND lokasi_id = $2;

-- name: CountDataMonitoring :one
SELECT COUNT(1) AS all, 
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00:00' AND '09:59') AS morning,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '12:00:00' AND '14:59') AS noon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00:00' AND '17:59') AS afternoon,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00:00' AND '23:59') AS night,
COUNT(1) FILTER (WHERE dibuat_pada::TIME BETWEEN '00:00:00' AND '05:59') AS midnight
FROM monitoring_data WHERE monitoring_terdaftar = $1;

-- name: AverageDataMonitoring :one
SELECT COALESCE(AVG(value), 0)::FLOAT AS all,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '06:00' AND '10:59'), 0)::FLOAT AS morning,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '10:00' AND '14:59'), 0)::FLOAT AS noon,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '15:00' AND '17:59'), 0)::FLOAT AS afternoon,
COALESCE(AVG(value) FILTER (WHERE dibuat_pada::TIME BETWEEN '18:00' AND '05:59'), 0)::FLOAT AS night
FROM monitoring_data WHERE monitoring_terdaftar = $1;