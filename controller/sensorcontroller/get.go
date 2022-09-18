package sensorcontroller

import (
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetAll(ctx *gin.Context) {
	sensors, err := ctr.service.GetAllTipeSensor(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensors,
	})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tipeSensor,
	})
}
