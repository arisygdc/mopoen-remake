package pgservice

import (
	"context"
	"mopoen-remake/database/postgres"
	"mopoen-remake/service/servicemodel"

	"github.com/google/uuid"
)

func (db postgre) DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error {
	param := postgres.CreateMonitoringTerdaftarParams{
		ID:           uuid.New(),
		TipeSensorID: daftarMonitoringParam.TipeSensor,
		LokasiID:     daftarMonitoringParam.Location_id,
		Nama:         daftarMonitoringParam.Nama,
		Keterangan:   daftarMonitoringParam.Keterangan,
	}

	return db.Queries.CreateMonitoringTerdaftar(ctx, param)
}

func (db postgre) CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64) error {
	param := postgres.CreateMonitoringValueParams{
		MonitoringTerdaftar: monitoring_id,
		Value:               value,
	}

	return db.Queries.CreateMonitoringValue(ctx, param)
}

func (db postgre) GetMonitoringTerdaftarByLokasi(ctx context.Context, lokasi_id int32) ([]servicemodel.MonitoringTerdaftar, error) {
	mtd, err := db.Queries.GetMonitoringTerdaftarByLokasi(ctx, lokasi_id)
	if err != nil {
		return nil, err
	}

	converted := make([]servicemodel.MonitoringTerdaftar, len(mtd))
	for i, v := range mtd {
		converted[i] = servicemodel.MonitoringTerdaftar(v)
	}

	return converted, nil
}

func (db postgre) GetMonitoringTerdaftar(ctx context.Context, id string) (servicemodel.DetailMonitoringTerdaftar, error) {
	idMon, err := uuid.Parse(id)
	monTdServiceModel := servicemodel.DetailMonitoringTerdaftar{}
	if err != nil {
		return servicemodel.DetailMonitoringTerdaftar{}, err
	}

	monTd, err := db.Queries.GetMonitoringTerdaftar(ctx, idMon)
	if err != nil {
		return monTdServiceModel, err
	}

	tipeSensor, err := db.Queries.GetTipeSensor(ctx, monTd.LokasiID)
	if err != nil {
		return monTdServiceModel, err
	}

	lokasi, err := db.Queries.FetchLokasi(ctx, monTd.LokasiID)
	if err != nil {
		return monTdServiceModel, err
	}

	monTdServiceModel = servicemodel.DetailMonitoringTerdaftar{
		ID:         monTd.ID,
		TipeSensor: servicemodel.TipeSensor(tipeSensor),
		LokasiID:   servicemodel.FetchLokasi(lokasi),
		Nama:       monTd.Nama,
		Keterangan: monTd.Keterangan,
	}

	return monTdServiceModel, err
}

func (db postgre) GetMonitoringData(ctx context.Context, id string) ([]servicemodel.MonitoringData, error) {
	idMon, err := uuid.Parse(id)
	var monData []servicemodel.MonitoringData
	if err != nil {
		return monData, err
	}
	row, err := db.Queries.GetMonitoringData(ctx, idMon)
	if err != nil {
		return monData, err
	}

	monData = make([]servicemodel.MonitoringData, len(row))
	for i, v := range row {
		monData[i] = servicemodel.MonitoringData(v)
	}
	return monData, err
}

func (db postgre) GetMonTerdaftarFilterLokasiAndSensor(ctx context.Context, lokasi_id int32, sensor_id int32) ([]servicemodel.MonitoringTerdaftar, error) {
	rows, err := db.Queries.GetMonTerdaftarFilterLokAndSensor(ctx, postgres.GetMonTerdaftarFilterLokAndSensorParams{
		LokasiID:     lokasi_id,
		TipeSensorID: sensor_id,
	})

	if err != nil {
		return nil, err
	}

	convert := make([]servicemodel.MonitoringTerdaftar, len(rows))
	for i, v := range rows {
		convert[i] = servicemodel.MonitoringTerdaftar(v)
	}
	return convert, err
}

func (db postgre) GetAnalisa(ctx context.Context, id uuid.UUID) (servicemodel.AnalisaMonitoring, error) {
	var analisa servicemodel.AnalisaMonitoring
	total, err := db.Queries.CountDataMonitoring(ctx, id)
	if err != nil {
		return servicemodel.AnalisaMonitoring{}, err
	}

	// this query return error when filter value is null
	average, _ := db.Queries.AverageDataMonitoring(ctx, id)
	analisa = servicemodel.AnalisaMonitoring{
		Overall: servicemodel.ResultMonitoring{
			Total:   total.All,
			Average: average.All,
		},
		Morning: servicemodel.ResultMonitoring{
			Total:   total.Morning,
			Average: average.Morning,
		},
		Afternoon: servicemodel.ResultMonitoring{
			Total:   total.Afternoon,
			Average: average.Afternoon,
		},
		Noon: servicemodel.ResultMonitoring{
			Total:   total.Noon,
			Average: average.Noon,
		},
		Night: servicemodel.ResultMonitoring{
			Total:   total.Night,
			Average: average.Night,
		},
		Midnight: servicemodel.ResultMonitoring{
			Total:   total.Midnight,
			Average: average.Midnight,
		},
	}
	return analisa, nil
}
