package monitoringcontroller

import (
	"log"
	"mopoen-remake/controller/request"
	"mopoen-remake/service/servicemodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) CreateDaftar(ctx *gin.Context) {
	req := request.PostDaftarMonitoring{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.DaftarMonitoring(ctx, servicemodel.DaftarMonitoring(req)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": req.Nama + " created",
	})
}

func (ctr Controller) CreateValue(ctx *gin.Context) {
	req := request.PostMonitoringValue{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.CreateMonitoringValue(ctx, req.KodeMonitoring, req.Value); err != nil {
		log.Printf("monitoring %v gagal memasukkan data", req.KodeMonitoring)
	}
}
