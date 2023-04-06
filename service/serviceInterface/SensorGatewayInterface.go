package serviceinterface

import (
	"context"

	"github.com/google/uuid"
)

type SensorGatewayInterface interface {
	// param: context context.Context, monitoring_id uuid.UUID, value float64
	CreateMonitoringValue(context.Context, uuid.UUID, float64) error
}
