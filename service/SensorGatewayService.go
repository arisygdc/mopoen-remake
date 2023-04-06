package service

import (
	"context"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"

	"github.com/google/uuid"
)

type SensorGatewayService struct {
	repo repository.Repository
}

func NewSensorGatewayService(repo repository.Repository) SensorGatewayService {
	return SensorGatewayService{repo: repo}
}

func (sgs SensorGatewayService) CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64) error {
	param := postgres.CreateMonitoringValueParams{
		MonitoringTerdaftar: monitoring_id,
		Value:               value,
	}

	return sgs.repo.CreateMonitoringValue(ctx, param)
}
