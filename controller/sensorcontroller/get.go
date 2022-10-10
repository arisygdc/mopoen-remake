package sensorcontroller

import (
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetAll(ctx *gin.Context) {
	sensors, err := ctr.service.GetAllTipeSensor(ctx)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespStatusOk(ctx, sensors)
}

func (ctr Controller) Get(ctx *gin.Context) {
	var idSensor request.GetSensor
	if err := ctx.ShouldBindUri(&idSensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tipeSensor, err := ctr.service.GetTipeSensor(ctx, idSensor.Id)
	if err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	helper.RespStatusOk(ctx, tipeSensor)
}
