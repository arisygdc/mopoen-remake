package sensorcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetAll(ctx *gin.Context) {
	sensors, err := ctr.service.GetAllTipeSensor(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensors,
	})
}

func (ctr Controller) Get(ctx *gin.Context) {
	var idSensor int32
	if err := ctx.ShouldBindQuery(&idSensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	sensors, err := ctr.service.GetAllTipeSensor(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensors,
	})
}
