-- name: CreateProvinsi :exec
INSERT INTO provinsi (nama) VALUES ($1);

-- name: CreateKabupaten :exec
INSERT INTO kabupaten (provinsi_id, nama) VALUES ($1, $2);

-- name: CreateKecamatan :exec
INSERT INTO kecamatan (kabupaten_id, nama) VALUES ($1, $2);

-- name: CreateDesa :exec
INSERT INTO desa (kecamatan_id, nama) VALUES ($1, $2);

-- name: DeleteProvinsi :one
DELETE FROM provinsi WHERE id = $1 RETURNING nama;

-- name: DeleteKabupaten :one
DELETE FROM kabupaten WHERE id = $1 RETURNING nama;

-- name: DeleteKecamatan :one
DELETE FROM kecamatan WHERE id = $1 RETURNING nama;

-- name: DeleteDesa :one
DELETE FROM desa WHERE id = $1 RETURNING nama;

-- name: GetAllProvinsi :many
SELECT * FROM provinsi;

-- name: GetAllKabupaten :many
SELECT * FROM kabupaten;

-- name: GetAllKecamatan :many
SELECT * FROM kecamatan;

-- name: GetAllDesa :many
SELECT * FROM desa;

-- name: GetKabupatenBy :many
SELECT * FROM kabupaten WHERE provinsi_id = $1;

-- name: GetKecamatanBy :many
SELECT * FROM kecamatan WHERE kabupaten_id = $1;

-- name: GetDesaBy :many
SELECT * FROM desa WHERE kecamatan_id = $1;

-- name: FetchLokasi :one
SELECT d.nama AS desa, kc.nama AS kecamatan, kb.nama AS kabupaten, p.nama AS provinsi
FROM desa d
INNER JOIN kecamatan kc ON d.kecamatan_id = kc.id
INNER JOIN kabupaten kb ON kc.kabupaten_id = kb.id
INNER JOIN provinsi p ON kb.provinsi_id = p.id
WHERE d.id = $1;