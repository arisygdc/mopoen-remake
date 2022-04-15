package repository

import (
	"context"
	"errors"
	"mopoen-remake/config"
	pgsvc "mopoen-remake/service/pg_service"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
)

var ErrDBDriverNotFound error = errors.New("database driver not found")

type IServices interface {
	CreateTipeSensor(ctx context.Context, tipe string, satuan string) error
	GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error)
	GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error)
	DeleteTipeSensor(ctx context.Context, id int32) error
	CreateProvinsi(ctx context.Context, provinsi string) error
	CreateKabupaten(ctx context.Context, provinsi_id int32, kabupaten string) error
	CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error
	CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error
	DeleteProvinsi(ctx context.Context, provinsi_id int32) error
	DeleteKabupaten(ctx context.Context, kabupaten_id int32) error
	DeleteKecamatan(ctx context.Context, kecamatan_id int32) error
	DeleteDesa(ctx context.Context, desa_id int32) error
	DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error
	CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64) error
}

func New(env config.Environment) (IServices, error) {
	if env.DBDriver == "postgres" {
		return pgsvc.NewPostgres(env)
	}
	return nil, ErrDBDriverNotFound
}
