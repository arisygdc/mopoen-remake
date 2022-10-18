package monitoringcontroller

import (
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	"mopoen-remake/service/servicemodel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctr Controller) CreateDaftar(ctx *gin.Context) {
	req := request.PostDaftarMonitoring{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	if err := ctr.service.DaftarMonitoring(ctx, servicemodel.DaftarMonitoring(req)); err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	helper.RespStatusOkWithMessage(ctx, req.Nama+" created")
}

func (ctr Controller) CreateValue(ctx *gin.Context) {
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
