package service

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"mopoen-remake/pkg/mail"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"
	"mopoen-remake/service/servicemodel"
	"os"
	"time"

	"mopoen-remake/pkg/utility"

	"github.com/google/uuid"
)

type MonitoringService struct {
	repo       repository.Repository
	mailSender mail.MailSender
}

func NewMonitoringService(repo repository.Repository, mailSender mail.MailSender) MonitoringService {
	return MonitoringService{
		repo:       repo,
		mailSender: mailSender,
	}
}

func (ls MonitoringService) DaftarMonitoring(ctx context.Context, daftarMonitoringParam servicemodel.DaftarMonitoring) error {
	param := postgres.CreateMonitoringTerdaftarParams{
		ID:           uuid.New(),
		TipeSensorID: daftarMonitoringParam.TipeSensor_id,
		LokasiID:     daftarMonitoringParam.Location_id,
		Email:        daftarMonitoringParam.Email,
		Author:       daftarMonitoringParam.Author,
		Nama:         daftarMonitoringParam.Nama,
		Keterangan:   daftarMonitoringParam.Keterangan,
	}

	// TODO
	// Transaction between create monitoring and send email
	// Asynchronous
	key := utility.HKDF16(param.ID.String(), param.Email, param.Author)
	param.Secret = hex.EncodeToString(key)
	created, err := ls.repo.CreateMonitoringTerdaftar(ctx, param)
	if err != nil {
		return err
	}
	return ls.mailSender.SendRegisteredMonitoring(created.Email, created.ID, created.Author, param.Secret)
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
		TipeSensor:   monTd.TipeSensor,
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

func (ls MonitoringService) GetCsvBuffer(ctx context.Context, id uuid.UUID) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := ls.EncodeToCsv(ctx, id, &buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func (ls MonitoringService) EncodeToCsv(ctx context.Context, id uuid.UUID, writer io.Writer) error {
	row, err := ls.repo.GetMonitoringData(ctx, id)
	if err != nil {
		return err
	}

	w := csv.NewWriter(writer)
	w.Write([]string{"Value", "Time"})
	for _, v := range row {
		value := fmt.Sprintf("%f", v.Value)
		w.Write([]string{value, v.DibuatPada.GoString()})
	}
	w.Flush()
	return nil
}

func (ls MonitoringService) SaveToCSV(ctx context.Context, id uuid.UUID) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s/generated_files/%s-%s.csv", pwd, id.String(), time.Now().Format("2006-01-02-15-04-05"))
	file, err := os.Open(filename)

	if err != nil {
		if err == os.ErrPermission {
			panic(err)
		}

		file, err = os.Create(filename)
		if err != nil {
			return "", err
		}

		ls.EncodeToCsv(ctx, id, file)
	}

	defer file.Close()

	return filename, nil
}
