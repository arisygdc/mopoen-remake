package repository

import (
	"context"
	"errors"
	"mopoen-remake/config"
	pgsvc "mopoen-remake/service/pg_service"
)

var ErrDBDriverNotFound error = errors.New("database driver not found")

type IServices interface {
	CreateTipeSensor(ctx context.Context, tipe string, satuan string) error
}

func New(env config.Environment) (IServices, error) {
	if env.DBDriver == "postgres" {
		return pgsvc.NewPostgres(env)
	}
	return nil, ErrDBDriverNotFound
}
