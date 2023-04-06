package service

import (
	"context"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"
	"mopoen-remake/service/servicemodel"
)

type SensorService struct {
	repo repository.Repository
}

func NewSensorService(repo repository.Repository) SensorService {
	return SensorService{repo: repo}
}

func (db SensorService) CreateTipeSensor(ctx context.Context, tipe string, satuan string) error {
	param := postgres.InsertTipeSensorParams{Tipe: tipe, Satuan: satuan}
	return db.repo.InsertTipeSensor(ctx, param)
}

func (db SensorService) GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error) {
	sensor, err := db.repo.GetTipeSensor(ctx, id)
	return servicemodel.TipeSensor(sensor), err
}

func (db SensorService) GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error) {
	sensors, err := db.repo.GetTipeSensors(ctx)
	if err != nil {
		return nil, err
	}

	sensorsConverted := make([]servicemodel.TipeSensor, len(sensors))
	for i, v := range sensors {
		sensorsConverted[i] = servicemodel.TipeSensor(v)
	}

	return sensorsConverted, nil
}

func (db SensorService) DeleteTipeSensor(ctx context.Context, id int32) error {
	return db.repo.DeleteTipeSensor(ctx, id)
}
