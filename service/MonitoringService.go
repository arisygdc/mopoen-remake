package service

import (
	"context"
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"
	"mopoen-remake/service/servicemodel"
	"os"

	"github.com/google/uuid"
)

type MonitoringService struct {
	repo repository.Repository
}

func NewMonitoringService(repo repository.Repository) MonitoringService {
	return MonitoringService{
		repo: repo,
	}
}

func (ls MonitoringService) DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error {
	param := postgres.CreateMonitoringTerdaftarParams{
		ID:           uuid.New(),
		TipeSensorID: daftarMonitoringParam.TipeSensor,
		LokasiID:     daftarMonitoringParam.Location_id,
		Nama:         daftarMonitoringParam.Nama,
		Keterangan:   daftarMonitoringParam.Keterangan,
	}

	return ls.repo.CreateMonitoringTerdaftar(ctx, param)
}

func (ls MonitoringService) GetMonitoringTerdaftar(ctx context.Context, option *servicemodel.GetMonitoringTerdaftarFilterOptions) ([]servicemodel.DetailMonitoringTerdaftar, error) {
	var converted []servicemodel.DetailMonitoringTerdaftar
	if option != nil {
		mtd, err := ls.repo.GetMonitoringTerdaftarFilter(ctx,
			postgres.GetMonitoringTerdaftarFilterParams{
				// tipe sensor
				Column1: option.TipeSensorID,
				Column2: option.LokasiID,
			})

		if err != nil {
			return converted, err
		}
		converted = make([]servicemodel.DetailMonitoringTerdaftar, len(mtd))
		for v := range mtd {
			converted[v] = servicemodel.DetailMonitoringTerdaftar{
				MonitoringID: mtd[v].MonitoringID,
				TipeSensorID: mtd[v].TipeSensorID,
				TipeSensor:   mtd[v].TipeSensor,
				Nama:         mtd[v].Nama,
				Keterangan:   mtd[v].Keterangan,
				Address:      mtd[v].Provinsi.String + ", " + mtd[v].Kabupaten.String + ", " + mtd[v].Kecamatan.String + ", " + mtd[v].Desa.String,
			}
		}
	}
	mtd, err := ls.repo.GetAllMonitoringTerdaftar(ctx)
	if err != nil {
		return converted, err
	}
	converted = make([]servicemodel.DetailMonitoringTerdaftar, len(mtd))
	for v := range mtd {
		converted[v] = servicemodel.DetailMonitoringTerdaftar{
			MonitoringID: mtd[v].MonitoringID,
			TipeSensorID: mtd[v].TipeSensorID,
			TipeSensor:   mtd[v].TipeSensor,
			Nama:         mtd[v].Nama,
			Keterangan:   mtd[v].Keterangan,
			Address:      mtd[v].Provinsi.String + ", " + mtd[v].Kabupaten.String + ", " + mtd[v].Kecamatan.String + ", " + mtd[v].Desa.String,
		}
	}
	return converted, nil
}

func (ls MonitoringService) GetMonitoringTerdaftarByID(ctx context.Context, id uuid.UUID) (servicemodel.DetailMonitoringTerdaftar, error) {
	var monTdServiceModel servicemodel.DetailMonitoringTerdaftar

	monTd, err := ls.repo.GetMonitoringTerdaftar(ctx, id)
	if err != nil {
		return monTdServiceModel, err
	}

	monTdServiceModel = servicemodel.DetailMonitoringTerdaftar{
		MonitoringID: monTd.MonitoringID,
		TipeSensorID: monTd.TipeSensorID,
		Nama:         monTd.Nama,
		Keterangan:   monTd.Keterangan,
		Address:      monTd.Address,
	}

	return monTdServiceModel, err
}

func (ls MonitoringService) GetMonitoringData(ctx context.Context, id string) ([]servicemodel.MonitoringData, error) {
	idMon, err := uuid.Parse(id)
	var monData []servicemodel.MonitoringData
	if err != nil {
		return monData, err
	}
	row, err := ls.repo.GetMonitoringData(ctx, idMon)
	if err != nil {
		return monData, err
	}

	monData = make([]servicemodel.MonitoringData, len(row))
	for i, v := range row {
		monData[i] = servicemodel.MonitoringData(v)
	}
	return monData, err
}

func (ls MonitoringService) GetMonTerdaftarFilterLokasiAndSensor(ctx context.Context, lokasi_id int32, sensor_id int32) ([]servicemodel.MonitoringTerdaftar, error) {
	rows, err := ls.repo.GetMonTerdaftarFilterLokAndSensor(ctx, postgres.GetMonTerdaftarFilterLokAndSensorParams{
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

func (ls MonitoringService) GetAnalisa(ctx context.Context, id uuid.UUID) (servicemodel.AnalisaMonitoring, error) {
	var analisa servicemodel.AnalisaMonitoring
	total, err := ls.repo.CountDataMonitoring(ctx, id)
	if err != nil {
		return analisa, err
	}

	average, err := ls.repo.AverageDataMonitoring(ctx, id)
	if err != nil {
		return analisa, err
	}

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
	}
	return analisa, nil
}

func (ls MonitoringService) ExtractToCSV(ctx context.Context, id uuid.UUID) (string, error) {
	row, err := ls.repo.GetMonitoringData(ctx, id)

	if err != nil {
		return "", err
	}

	hParam1 := []byte(row[len(row)].DibuatPada.String())
	hParam2 := []byte(row[0].DibuatPada.String())

	hParam := append(hParam1, hParam2...)

	hClass := sha1.New()
	hClass.Write(hParam)

	hSum := hClass.Sum(nil)

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s/generated_files/%s-%s.csv", pwd, id.String(), string(hSum))
	file, err := os.Open(filename)

	switch err {
	case os.ErrNotExist:
		file, err = os.Create(filename)
		if err != nil {
			return "", err
		}

		writer := csv.NewWriter(file)

		for _, v := range row {
			value := fmt.Sprintf("%f", v.Value)
			convert := []string{v.DibuatPada.String(), value}
			writer.Write(convert)
		}
		writer.Flush()
	case os.ErrPermission:
		panic(err)
	}

	defer file.Close()

	return filename, nil
}
