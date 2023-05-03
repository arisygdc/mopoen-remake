package serviceinterface

import (
	"context"

	"github.com/google/uuid"
)

type SensorGatewayInterface interface {
	// param: context context.Context, monitoring_id uuid.UUID, value float64 secret string
	CreateMonitoringValue(context.Context, uuid.UUID, float64, string) error
}
