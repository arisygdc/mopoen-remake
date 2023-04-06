package sensorgateway

import (
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	svc "mopoen-remake/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SensorGatewayController struct {
	service svc.IServices
}

func New(service svc.IServices) SensorGatewayController {
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

	if err := ctr.service.CreateMonitoringValue(ctx, id, req.Value); err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}
