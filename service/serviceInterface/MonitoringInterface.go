package serviceinterface

import (
	"bytes"
	"context"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
)

type MonitoringInterface interface {
	// param: context context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring
	DaftarMonitoring(context.Context, servicemodel.DaftarMonitoring) error
	// param: context context.Context, filter_option lokasi_id int32 || tipe_sensor_id int32
	GetMonitoringTerdaftar(context.Context, *servicemodel.GetMonitoringTerdaftarFilterOptions) ([]servicemodel.DetailMonitoringTerdaftar, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// return: information about monitoring
	GetMonitoringTerdaftarByID(context.Context, uuid.UUID) (servicemodel.DetailMonitoringTerdaftar, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// return: List value monitoring
	GetMonitoringData(context.Context, string) ([]servicemodel.MonitoringData, error)
	// param: context context.Context, monitoring_id uuid.UUID
	// count average value monitoring per sub day
	// return: AnalisaMonitoring
	GetAnalisa(context.Context, uuid.UUID) (servicemodel.AnalisaMonitoring, error)
	// param: context context.Context, id uuid.UUID
	// extract monitoring data to csv and save to disk
	// return: string path file
	SaveToCSV(context.Context, uuid.UUID) (string, error)
	// param: context context.Context, id uuid.UUID
	// extract monitoring data to csv
	// return as buffer wihout save to disk
	GetCsvBuffer(context.Context, uuid.UUID) (*bytes.Buffer, error)
}
