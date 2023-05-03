package sensorgateway

import (
	"mopoen-remake/controller/helper"
	"mopoen-remake/request"
	ifSG "mopoen-remake/service/serviceInterface"
	"mopoen-remake/service/servicemodel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SensorGatewayController struct {
	service ifSG.SensorGatewayInterface
}

func NewSensorGatewayController(service ifSG.SensorGatewayInterface) SensorGatewayController {
	return SensorGatewayController{service: service}
}

// SaveDataFromSensor responsible to save data from sensor to database
func (ctr SensorGatewayController) SaveDataFromSensor(ctx *gin.Context) {
	req := request.PostMonitoringValue{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	id, err := uuid.Parse(req.KodeMonitoring)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	if err := ctr.service.CreateMonitoringValue(ctx, id, req.Value, req.Secret); err != nil {
		if err == servicemodel.ErrWrongSecret {
			helper.RespBadRequest(ctx, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}
