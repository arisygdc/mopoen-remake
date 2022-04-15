package sensorcontroller

import (
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) Delete(ctx *gin.Context) {
	req := request.DeleteSensor{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctr.service.DeleteTipeSensor(ctx, req.Id)
}
