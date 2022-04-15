-- name: CreateProvinsi :exec
INSERT INTO provinsi (nama) VALUES ($1);

-- name: CreateKabupaten :exec
INSERT INTO kabupaten (provinsi_id, nama) VALUES ($1, $2);

-- name: CreateKecamatan :exec
INSERT INTO kecamatan (kabupaten_id, nama) VALUES ($1, $2);

-- name: CreateDesa :exec
INSERT INTO desa (kecamatan_id, nama) VALUES ($1, $2);

-- name: DeleteProvinsi :exec
DELETE FROM provinsi WHERE id = $1;

-- name: DeleteKabupaten :exec
DELETE FROM kabupaten WHERE id = $1;

-- name: DeleteKecamatan :exec
DELETE FROM kecamatan WHERE id = $1;

-- name: DeleteDesa :exec
DELETE FROM desa WHERE id = $1;