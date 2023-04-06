-- name: InsertTipeSensor :exec
INSERT INTO tipe_sensor (tipe, satuan) VALUES ($1, $2);

-- name: DeleteTipeSensor :exec
DELETE FROM tipe_sensor WHERE id = $1;

-- name: GetTipeSensors :many
SELECT * FROM tipe_sensor;

-- name: GetTipeSensor :one
SELECT * FROM tipe_sensor WHERE id = $1;