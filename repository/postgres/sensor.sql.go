// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: sensor.sql

package postgres

import (
	"context"
)

const deleteTipeSensor = `-- name: DeleteTipeSensor :exec
DELETE FROM tipe_sensor WHERE id = $1
`

func (q *Queries) DeleteTipeSensor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTipeSensor, id)
	return err
}

const getTipeSensor = `-- name: GetTipeSensor :one
SELECT id, tipe, satuan FROM tipe_sensor WHERE id = $1
`

func (q *Queries) GetTipeSensor(ctx context.Context, id int32) (TipeSensor, error) {
	row := q.db.QueryRowContext(ctx, getTipeSensor, id)
	var i TipeSensor
	err := row.Scan(&i.ID, &i.Tipe, &i.Satuan)
	return i, err
}

const getTipeSensors = `-- name: GetTipeSensors :many
SELECT id, tipe, satuan FROM tipe_sensor
`

func (q *Queries) GetTipeSensors(ctx context.Context) ([]TipeSensor, error) {
	rows, err := q.db.QueryContext(ctx, getTipeSensors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TipeSensor
	for rows.Next() {
		var i TipeSensor
		if err := rows.Scan(&i.ID, &i.Tipe, &i.Satuan); err != nil {
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

const insertTipeSensor = `-- name: InsertTipeSensor :exec
INSERT INTO tipe_sensor (tipe, satuan) VALUES ($1, $2)
`

type InsertTipeSensorParams struct {
	Tipe   string `json:"tipe"`
	Satuan string `json:"satuan"`
}

func (q *Queries) InsertTipeSensor(ctx context.Context, arg InsertTipeSensorParams) error {
	_, err := q.db.ExecContext(ctx, insertTipeSensor, arg.Tipe, arg.Satuan)
	return err
}