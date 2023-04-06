package serviceinterface

import (
	"context"
	"mopoen-remake/service/servicemodel"
)

type SensorInterface interface {
	// param: context context.Context, tipe string, satuan string
	CreateTipeSensor(ctx context.Context, tipe string, satuan string) error
	// param: context context.Context, tipe_sensor_id int32
	GetTipeSensor(ctx context.Context, id int32) (servicemodel.TipeSensor, error)
	// param: context context.Context, tipe_sensor_id int32
	DeleteTipeSensor(ctx context.Context, id int32) error
	// return list of tipe sensor
	GetAllTipeSensor(ctx context.Context) ([]servicemodel.TipeSensor, error)
}
