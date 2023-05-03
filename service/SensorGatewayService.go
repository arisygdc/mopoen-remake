package service

import (
	"context"
	"mopoen-remake/repository"
	"mopoen-remake/repository/postgres"
	"mopoen-remake/service/servicemodel"
	"strings"

	"github.com/google/uuid"
)

type SensorGatewayService struct {
	repo repository.Repository
}

func NewSensorGatewayService(repo repository.Repository) SensorGatewayService {
	return SensorGatewayService{repo: repo}
}

func (sgs SensorGatewayService) CreateMonitoringValue(ctx context.Context, monitoring_id uuid.UUID, value float64, secret string) error {
	param := postgres.CreateMonitoringValueParams{
		ID:     monitoring_id,
		Value:  value,
		Secret: secret,
	}
	err := sgs.repo.CreateMonitoringValue(ctx, param)
	if err != nil {
		strErr := err.Error()
		if strings.Contains(strErr, "violates not-null constraint") {
			return servicemodel.ErrWrongSecret
		}
		return err
	}
	return nil
}
