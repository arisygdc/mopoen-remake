package repository

import (
	"context"
	"errors"
	"mopoen-remake/config"
	pgservice "mopoen-remake/service/pg_service"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
)

const (
	PostgreDriver = "postgres"
)

var ErrDBDriverNotFound = errors.New("database driver not found")

type IServices interface {
	CreateTipeSensor(ctx context.Context, tipe string, satuan string) error
	GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error)
	GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error)
	DeleteTipeSensor(ctx context.Context, id int32) error
	CreateProvinsi(ctx context.Context, provinsi string) error
	CreateKabupaten(ctx context.Context, provinsi_id int32, kabupaten string) error
	CreateKecamatan(ctx context.Context, kabupaten_id int32, kecamatan string) error
	CreateDesa(ctx context.Context, kecamatan_id int32, desa string) error
	DeleteProvinsi(ctx context.Context, provinsi_id int32) (string, error)
	DeleteKabupaten(ctx context.Context, kabupaten_id int32) (string, error)
	DeleteKecamatan(ctx context.Context, kecamatan_id int32) (string, error)
	DeleteDesa(ctx context.Context, desa_id int32) (string, error)
	DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error
	CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64) error
	GetAllProvinsi(ctx context.Context) ([]servicemodel.Provinsi, error)
	GetAllKabupaten(ctx context.Context) ([]servicemodel.Kabupaten, error)
	GetAllKecamatan(ctx context.Context) ([]servicemodel.Kecamatan, error)
	GetAllDesa(ctx context.Context) ([]servicemodel.Desa, error)
	GetLokasiBy(ctx context.Context, tipe string, depends int32) ([]servicemodel.Lokasi, error)
	GetMonitoringTerdaftar(ctx context.Context, id string) (servicemodel.DetailMonitoringTerdaftar, error)
	GetMonitoringTerdaftarByLokasi(ctx context.Context, lokasi_id int32) ([]servicemodel.MonitoringTerdaftar, error)
	GetMonitoringData(ctx context.Context, id string) ([]servicemodel.MonitoringData, error)
	GetMonTerdaftarFilterLokasiAndSensor(ctx context.Context, lokasi_id int32, sensor_id int32) ([]servicemodel.MonitoringTerdaftar, error)
	GetAnalisa(ctx context.Context, id uuid.UUID) (servicemodel.AnalisaMonitoring, error)
}

func New(env config.Environment) (IServices, error) {
	if env.DBDriver == PostgreDriver {
		return pgservice.NewPostgres(env.DBDriver, env.DBSource)
	}
	return nil, ErrDBDriverNotFound
}
