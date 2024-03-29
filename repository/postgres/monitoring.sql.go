// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: monitoring.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const averageDataMonitoring = `-- name: AverageDataMonitoring :one
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
    monitoring_terdaftar = $1
`

type AverageDataMonitoringRow struct {
	All       float64 `json:"all"`
	Morning   float64 `json:"morning"`
	Noon      float64 `json:"noon"`
	Afternoon float64 `json:"afternoon"`
	Night     float64 `json:"night"`
}

func (q *Queries) AverageDataMonitoring(ctx context.Context, monitoringTerdaftar uuid.UUID) (AverageDataMonitoringRow, error) {
	row := q.db.QueryRowContext(ctx, averageDataMonitoring, monitoringTerdaftar)
	var i AverageDataMonitoringRow
	err := row.Scan(
		&i.All,
		&i.Morning,
		&i.Noon,
		&i.Afternoon,
		&i.Night,
	)
	return i, err
}

const countDataMonitoring = `-- name: CountDataMonitoring :one
SELECT 
    COUNT(1) AS all, 
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 6 AND 10) AS morning,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 11 AND 14) AS noon,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) BETWEEN 15 AND 17) AS afternoon,
    COUNT(1) FILTER (WHERE EXTRACT(HOUR FROM dibuat_pada) >= 18 OR EXTRACT(HOUR FROM dibuat_pada) < 6) AS night
FROM 
    monitoring_data
WHERE 
    monitoring_terdaftar = $1
`

type CountDataMonitoringRow struct {
	All       int64 `json:"all"`
	Morning   int64 `json:"morning"`
	Noon      int64 `json:"noon"`
	Afternoon int64 `json:"afternoon"`
	Night     int64 `json:"night"`
}

func (q *Queries) CountDataMonitoring(ctx context.Context, monitoringTerdaftar uuid.UUID) (CountDataMonitoringRow, error) {
	row := q.db.QueryRowContext(ctx, countDataMonitoring, monitoringTerdaftar)
	var i CountDataMonitoringRow
	err := row.Scan(
		&i.All,
		&i.Morning,
		&i.Noon,
		&i.Afternoon,
		&i.Night,
	)
	return i, err
}

const createMonitoringTerdaftar = `-- name: CreateMonitoringTerdaftar :one
INSERT INTO monitoring_terdaftar (id, tipe_sensor_id, lokasi_id, email, secret, author, nama, keterangan) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, tipe_sensor_id, lokasi_id, email, author, secret, nama, keterangan
`

type CreateMonitoringTerdaftarParams struct {
	ID           uuid.UUID `json:"id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	LokasiID     int32     `json:"lokasi_id"`
	Email        string    `json:"email"`
	Secret       string    `json:"secret"`
	Author       string    `json:"author"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
}

func (q *Queries) CreateMonitoringTerdaftar(ctx context.Context, arg CreateMonitoringTerdaftarParams) (MonitoringTerdaftar, error) {
	row := q.db.QueryRowContext(ctx, createMonitoringTerdaftar,
		arg.ID,
		arg.TipeSensorID,
		arg.LokasiID,
		arg.Email,
		arg.Secret,
		arg.Author,
		arg.Nama,
		arg.Keterangan,
	)
	var i MonitoringTerdaftar
	err := row.Scan(
		&i.ID,
		&i.TipeSensorID,
		&i.LokasiID,
		&i.Email,
		&i.Author,
		&i.Secret,
		&i.Nama,
		&i.Keterangan,
	)
	return i, err
}

const createMonitoringValue = `-- name: CreateMonitoringValue :exec
INSERT INTO monitoring_data (monitoring_terdaftar, value) VALUES 
(
    (
        SELECT monitoring_terdaftar.id 
        FROM monitoring_terdaftar 
        WHERE monitoring_terdaftar.id=$1 
        AND monitoring_terdaftar.secret=$3
    ), 
    $2
)
`

type CreateMonitoringValueParams struct {
	ID     uuid.UUID `json:"id"`
	Value  float64   `json:"value"`
	Secret string    `json:"secret"`
}

func (q *Queries) CreateMonitoringValue(ctx context.Context, arg CreateMonitoringValueParams) error {
	_, err := q.db.ExecContext(ctx, createMonitoringValue, arg.ID, arg.Value, arg.Secret)
	return err
}

const getAllMonitoringTerdaftar = `-- name: GetAllMonitoringTerdaftar :many
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
`

type GetAllMonitoringTerdaftarRow struct {
	MonitoringID uuid.UUID      `json:"monitoring_id"`
	TipeSensorID int32          `json:"tipe_sensor_id"`
	TipeSensor   string         `json:"tipe_sensor"`
	Nama         string         `json:"nama"`
	Keterangan   string         `json:"keterangan"`
	Provinsi     sql.NullString `json:"provinsi"`
	Kecamatan    sql.NullString `json:"kecamatan"`
	Kabupaten    sql.NullString `json:"kabupaten"`
	Desa         sql.NullString `json:"desa"`
}

func (q *Queries) GetAllMonitoringTerdaftar(ctx context.Context) ([]GetAllMonitoringTerdaftarRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllMonitoringTerdaftar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllMonitoringTerdaftarRow
	for rows.Next() {
		var i GetAllMonitoringTerdaftarRow
		if err := rows.Scan(
			&i.MonitoringID,
			&i.TipeSensorID,
			&i.TipeSensor,
			&i.Nama,
			&i.Keterangan,
			&i.Provinsi,
			&i.Kecamatan,
			&i.Kabupaten,
			&i.Desa,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonTerdaftarFilterLokAndSensor = `-- name: GetMonTerdaftarFilterLokAndSensor :many
SELECT id, tipe_sensor_id, lokasi_id, email, author, nama, keterangan FROM monitoring_terdaftar WHERE tipe_sensor_id = $1 AND lokasi_id = $2
`

type GetMonTerdaftarFilterLokAndSensorParams struct {
	TipeSensorID int32 `json:"tipe_sensor_id"`
	LokasiID     int32 `json:"lokasi_id"`
}

type GetMonTerdaftarFilterLokAndSensorRow struct {
	ID           uuid.UUID `json:"id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	LokasiID     int32     `json:"lokasi_id"`
	Email        string    `json:"email"`
	Author       string    `json:"author"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
}

func (q *Queries) GetMonTerdaftarFilterLokAndSensor(ctx context.Context, arg GetMonTerdaftarFilterLokAndSensorParams) ([]GetMonTerdaftarFilterLokAndSensorRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonTerdaftarFilterLokAndSensor, arg.TipeSensorID, arg.LokasiID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonTerdaftarFilterLokAndSensorRow
	for rows.Next() {
		var i GetMonTerdaftarFilterLokAndSensorRow
		if err := rows.Scan(
			&i.ID,
			&i.TipeSensorID,
			&i.LokasiID,
			&i.Email,
			&i.Author,
			&i.Nama,
			&i.Keterangan,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonitoringData = `-- name: GetMonitoringData :many
SELECT value, date(dibuat_pada)::varchar AS date, (dibuat_pada::time)::varchar AS time FROM monitoring_data WHERE monitoring_terdaftar = $1
`

type GetMonitoringDataRow struct {
	Value float64 `json:"value"`
	Date  string  `json:"date"`
	Time  string  `json:"time"`
}

func (q *Queries) GetMonitoringData(ctx context.Context, monitoringTerdaftar uuid.UUID) ([]GetMonitoringDataRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonitoringData, monitoringTerdaftar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonitoringDataRow
	for rows.Next() {
		var i GetMonitoringDataRow
		if err := rows.Scan(&i.Value, &i.Date, &i.Time); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonitoringTerdaftar = `-- name: GetMonitoringTerdaftar :one
SELECT 
    mt.id as monitoring_id, 
    mt.tipe_sensor_id, 
    ts.tipe::varchar,
    ts.satuan::varchar, 
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
    mt.id = $1
`

type GetMonitoringTerdaftarRow struct {
	MonitoringID uuid.UUID `json:"monitoring_id"`
	TipeSensorID int32     `json:"tipe_sensor_id"`
	TsTipe       string    `json:"ts_tipe"`
	TsSatuan     string    `json:"ts_satuan"`
	Nama         string    `json:"nama"`
	Keterangan   string    `json:"keterangan"`
	Address      string    `json:"address"`
}

func (q *Queries) GetMonitoringTerdaftar(ctx context.Context, id uuid.UUID) (GetMonitoringTerdaftarRow, error) {
	row := q.db.QueryRowContext(ctx, getMonitoringTerdaftar, id)
	var i GetMonitoringTerdaftarRow
	err := row.Scan(
		&i.MonitoringID,
		&i.TipeSensorID,
		&i.TsTipe,
		&i.TsSatuan,
		&i.Nama,
		&i.Keterangan,
		&i.Address,
	)
	return i, err
}

const getMonitoringTerdaftarFilter = `-- name: GetMonitoringTerdaftarFilter :many
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
END
`

type GetMonitoringTerdaftarFilterParams struct {
	Column1 interface{} `json:"column_1"`
	Column2 interface{} `json:"column_2"`
}

type GetMonitoringTerdaftarFilterRow struct {
	MonitoringID uuid.UUID      `json:"monitoring_id"`
	TipeSensorID int32          `json:"tipe_sensor_id"`
	TipeSensor   string         `json:"tipe_sensor"`
	Nama         string         `json:"nama"`
	Keterangan   string         `json:"keterangan"`
	Provinsi     sql.NullString `json:"provinsi"`
	Kecamatan    sql.NullString `json:"kecamatan"`
	Kabupaten    sql.NullString `json:"kabupaten"`
	Desa         sql.NullString `json:"desa"`
}

func (q *Queries) GetMonitoringTerdaftarFilter(ctx context.Context, arg GetMonitoringTerdaftarFilterParams) ([]GetMonitoringTerdaftarFilterRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonitoringTerdaftarFilter, arg.Column1, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMonitoringTerdaftarFilterRow
	for rows.Next() {
		var i GetMonitoringTerdaftarFilterRow
		if err := rows.Scan(
			&i.MonitoringID,
			&i.TipeSensorID,
			&i.TipeSensor,
			&i.Nama,
			&i.Keterangan,
			&i.Provinsi,
			&i.Kecamatan,
			&i.Kabupaten,
			&i.Desa,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
