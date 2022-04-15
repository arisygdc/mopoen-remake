package sensorcontroller

import (
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) Create(ctx *gin.Context) {
	req := request.PostSensor{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := ctr.service.CreateTipeSensor(ctx, req.Tipe, req.Satuan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
