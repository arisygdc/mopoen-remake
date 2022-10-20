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
	CreateTipeSensor(context.Context, string, string) error
	GetTipeSensor(context.Context, int32) (servicemodel.TipeSensor, error)
	GetAllTipeSensor(context.Context) ([]servicemodel.TipeSensor, error)
	DeleteTipeSensor(context.Context, int32) error
	CreateProvinsi(context.Context, string) error
	CreateKabupaten(context.Context, int32, string) error
	CreateKecamatan(context.Context, int32, string) error
	CreateDesa(context.Context, int32, string) error
	DeleteProvinsi(context.Context, int32) (string, error)
	DeleteKabupaten(context.Context, int32) (string, error)
	DeleteKecamatan(context.Context, int32) (string, error)
	DeleteDesa(context.Context, int32) (string, error)
	DaftarMonitoring(context.Context, servicemodel.DaftarMonitoring) error
	CreateMonitoringValue(context.Context, uuid.UUID, float64) error
	GetAllProvinsi(context.Context) ([]servicemodel.Provinsi, error)
	GetAllKabupaten(context.Context) ([]servicemodel.Kabupaten, error)
	GetAllKecamatan(context.Context) ([]servicemodel.Kecamatan, error)
	GetAllDesa(context.Context) ([]servicemodel.Desa, error)
	GetLokasiBy(context.Context, string, int32) ([]servicemodel.Lokasi, error)
	GetMonitoringTerdaftar(context.Context, string) (servicemodel.DetailMonitoringTerdaftar, error)
	GetMonitoringTerdaftarByLokasi(context.Context, int32) ([]servicemodel.MonitoringTerdaftar, error)
	GetMonitoringData(context.Context, string) ([]servicemodel.MonitoringData, error)
	GetMonTerdaftarFilterLokasiAndSensor(context.Context, int32, int32) ([]servicemodel.MonitoringTerdaftar, error)
	GetAnalisa(context.Context, uuid.UUID) (servicemodel.AnalisaMonitoring, error)
	ExtractToCSV(context.Context, uuid.UUID) (string, error)
}

func New(env config.Environment) (IServices, error) {
	if env.DBDriver == PostgreDriver {
		return pgservice.NewPostgres(env.DBDriver, env.DBSource)
	}
	return nil, ErrDBDriverNotFound
}
