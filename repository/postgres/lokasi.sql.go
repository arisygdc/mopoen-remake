// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: lokasi.sql

package postgres

import (
	"context"
)

const createDesa = `-- name: CreateDesa :exec
INSERT INTO desa (kecamatan_id, nama) VALUES ($1, $2)
`

type CreateDesaParams struct {
	KecamatanID int32  `json:"kecamatan_id"`
	Nama        string `json:"nama"`
}

func (q *Queries) CreateDesa(ctx context.Context, arg CreateDesaParams) error {
	_, err := q.db.ExecContext(ctx, createDesa, arg.KecamatanID, arg.Nama)
	return err
}

const createKabupaten = `-- name: CreateKabupaten :exec
INSERT INTO kabupaten (provinsi_id, nama) VALUES ($1, $2)
`

type CreateKabupatenParams struct {
	ProvinsiID int32  `json:"provinsi_id"`
	Nama       string `json:"nama"`
}

func (q *Queries) CreateKabupaten(ctx context.Context, arg CreateKabupatenParams) error {
	_, err := q.db.ExecContext(ctx, createKabupaten, arg.ProvinsiID, arg.Nama)
	return err
}

const createKecamatan = `-- name: CreateKecamatan :exec
INSERT INTO kecamatan (kabupaten_id, nama) VALUES ($1, $2)
`

type CreateKecamatanParams struct {
	KabupatenID int32  `json:"kabupaten_id"`
	Nama        string `json:"nama"`
}

func (q *Queries) CreateKecamatan(ctx context.Context, arg CreateKecamatanParams) error {
	_, err := q.db.ExecContext(ctx, createKecamatan, arg.KabupatenID, arg.Nama)
	return err
}

const createProvinsi = `-- name: CreateProvinsi :exec
INSERT INTO provinsi (nama) VALUES ($1)
`

func (q *Queries) CreateProvinsi(ctx context.Context, nama string) error {
	_, err := q.db.ExecContext(ctx, createProvinsi, nama)
	return err
}

const deleteDesa = `-- name: DeleteDesa :one
DELETE FROM desa WHERE id = $1 RETURNING nama
`

func (q *Queries) DeleteDesa(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteDesa, id)
	var nama string
	err := row.Scan(&nama)
	return nama, err
}

const deleteKabupaten = `-- name: DeleteKabupaten :one
DELETE FROM kabupaten WHERE id = $1 RETURNING nama
`

func (q *Queries) DeleteKabupaten(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteKabupaten, id)
	var nama string
	err := row.Scan(&nama)
	return nama, err
}

const deleteKecamatan = `-- name: DeleteKecamatan :one
DELETE FROM kecamatan WHERE id = $1 RETURNING nama
`

func (q *Queries) DeleteKecamatan(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteKecamatan, id)
	var nama string
	err := row.Scan(&nama)
	return nama, err
}

const deleteProvinsi = `-- name: DeleteProvinsi :one
DELETE FROM provinsi WHERE id = $1 RETURNING nama
`

func (q *Queries) DeleteProvinsi(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteProvinsi, id)
	var nama string
	err := row.Scan(&nama)
	return nama, err
}

const fetchLokasi = `-- name: FetchLokasi :one
SELECT d.nama AS desa, kc.nama AS kecamatan, kb.nama AS kabupaten, p.nama AS provinsi
FROM desa d
INNER JOIN kecamatan kc ON d.kecamatan_id = kc.id
INNER JOIN kabupaten kb ON kc.kabupaten_id = kb.id
INNER JOIN provinsi p ON kb.provinsi_id = p.id
WHERE d.id = $1
`

type FetchLokasiRow struct {
	Desa      string `json:"desa"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
}

func (q *Queries) FetchLokasi(ctx context.Context, id int32) (FetchLokasiRow, error) {
	row := q.db.QueryRowContext(ctx, fetchLokasi, id)
	var i FetchLokasiRow
	err := row.Scan(
		&i.Desa,
		&i.Kecamatan,
		&i.Kabupaten,
		&i.Provinsi,
	)
	return i, err
}

const getAllDesa = `-- name: GetAllDesa :many
SELECT id, kecamatan_id, nama FROM desa
`

func (q *Queries) GetAllDesa(ctx context.Context) ([]Desa, error) {
	rows, err := q.db.QueryContext(ctx, getAllDesa)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Desa
	for rows.Next() {
		var i Desa
		if err := rows.Scan(&i.ID, &i.KecamatanID, &i.Nama); err != nil {
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

const getAllKabupaten = `-- name: GetAllKabupaten :many
SELECT id, provinsi_id, nama FROM kabupaten
`

func (q *Queries) GetAllKabupaten(ctx context.Context) ([]Kabupaten, error) {
	rows, err := q.db.QueryContext(ctx, getAllKabupaten)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Kabupaten
	for rows.Next() {
		var i Kabupaten
		if err := rows.Scan(&i.ID, &i.ProvinsiID, &i.Nama); err != nil {
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

const getAllKecamatan = `-- name: GetAllKecamatan :many
SELECT id, kabupaten_id, nama FROM kecamatan
`

func (q *Queries) GetAllKecamatan(ctx context.Context) ([]Kecamatan, error) {
	rows, err := q.db.QueryContext(ctx, getAllKecamatan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Kecamatan
	for rows.Next() {
		var i Kecamatan
		if err := rows.Scan(&i.ID, &i.KabupatenID, &i.Nama); err != nil {
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

const getAllProvinsi = `-- name: GetAllProvinsi :many
SELECT id, nama FROM provinsi
`

func (q *Queries) GetAllProvinsi(ctx context.Context) ([]Provinsi, error) {
	rows, err := q.db.QueryContext(ctx, getAllProvinsi)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Provinsi
	for rows.Next() {
		var i Provinsi
		if err := rows.Scan(&i.ID, &i.Nama); err != nil {
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

const getDesaBy = `-- name: GetDesaBy :many
SELECT id, kecamatan_id, nama FROM desa WHERE kecamatan_id = $1
`

func (q *Queries) GetDesaBy(ctx context.Context, kecamatanID int32) ([]Desa, error) {
	rows, err := q.db.QueryContext(ctx, getDesaBy, kecamatanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Desa
	for rows.Next() {
		var i Desa
		if err := rows.Scan(&i.ID, &i.KecamatanID, &i.Nama); err != nil {
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

const getKabupatenBy = `-- name: GetKabupatenBy :many
SELECT id, provinsi_id, nama FROM kabupaten WHERE provinsi_id = $1
`

func (q *Queries) GetKabupatenBy(ctx context.Context, provinsiID int32) ([]Kabupaten, error) {
	rows, err := q.db.QueryContext(ctx, getKabupatenBy, provinsiID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Kabupaten
	for rows.Next() {
		var i Kabupaten
		if err := rows.Scan(&i.ID, &i.ProvinsiID, &i.Nama); err != nil {
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

const getKecamatanBy = `-- name: GetKecamatanBy :many
SELECT id, kabupaten_id, nama FROM kecamatan WHERE kabupaten_id = $1
`

func (q *Queries) GetKecamatanBy(ctx context.Context, kabupatenID int32) ([]Kecamatan, error) {
	rows, err := q.db.QueryContext(ctx, getKecamatanBy, kabupatenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Kecamatan
	for rows.Next() {
		var i Kecamatan
		if err := rows.Scan(&i.ID, &i.KabupatenID, &i.Nama); err != nil {
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
