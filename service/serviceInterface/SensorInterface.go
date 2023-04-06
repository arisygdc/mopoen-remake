package serviceinterface

import (
	"context"
	"mopoen-remake/service/servicemodel"
)

type SensorInterface interface {
	// param: context context.Context, tipe string, satuan string
	CreateTipeSensor(context.Context, string, string) error
	// param: context context.Context, tipe_sensor_id int32
	GetTipeSensor(context.Context, int32) (servicemodel.TipeSensor, error)
	// param: context context.Context, tipe_sensor_id int32
	DeleteTipeSensor(context.Context, int32) error
	// return list of tipe sensor
	GetAllTipeSensor(context.Context) ([]servicemodel.TipeSensor, error)
}
