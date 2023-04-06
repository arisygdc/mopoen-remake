package serviceinterface

import (
	"context"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
)

type MonitoringInterface interface {
	// param: context context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring
	DaftarMonitoring(context.Context, servicemodel.DaftarMonitoring) error
	// param: context context.Context, lokasi_id int32
	GetMonitoringTerdaftarByLokasi(context.Context, int32) ([]servicemodel.MonitoringTerdaftar, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// return: information about monitoring
	GetMonitoringTerdaftar(context.Context, string) (servicemodel.DetailMonitoringTerdaftar, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// return: List value monitoring
	GetMonitoringData(context.Context, string) ([]servicemodel.MonitoringData, error)
	// param: context context.Context, lokasi_id int32, sensor_id int32
	// exec same as GetMonitoringTerdaftarByLokasi but with filter sensor
	// return: information about monitoring
	GetMonTerdaftarFilterLokasiAndSensor(context.Context, int32, int32) ([]servicemodel.MonitoringTerdaftar, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// count average value monitoring per sub day
	// return: AnalisaMonitoring
	GetAnalisa(context.Context, uuid.UUID) (servicemodel.AnalisaMonitoring, error)
	// param: context context.Context, id uuid.UUID
	// extract monitoring data to csv file
	// return: string path file
	ExtractToCSV(context.Context, uuid.UUID) (string, error)
}
