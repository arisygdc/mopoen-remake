package monitoringcontroller

import (
	"errors"
	"io"
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ErrQParam = errors.New("query param tidak valid")

func (ctr Controller) GetTerdaftar(ctx *gin.Context) {
	QLok, okLok := ctx.GetQuery("lokasi_id")
	QSensor, okSensor := ctx.GetQuery("sensor_id")
	lok_id, _ := strconv.Atoi(QLok)
	sensor_id, _ := strconv.Atoi(QSensor)
	var err error

	if okLok && okSensor && lok_id > 0 && sensor_id > 0 {
		mtd, qErr := ctr.service.GetMonTerdaftarFilterLokasiAndSensor(ctx, int32(lok_id), int32(sensor_id))
		if qErr != nil {
			helper.RespBadRequest(ctx, qErr.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": mtd,
		})
		return
	}

	if okLok && lok_id > 0 {
		mtd, qErr := ctr.service.GetMonitoringTerdaftarByLokasi(ctx, int32(lok_id))
		if qErr != nil {
			helper.RespBadRequest(ctx, qErr.Error())
			return
		}

		helper.RespStatusOk(ctx, mtd)
		return
	}

	helper.RespBadRequest(ctx, err.Error())
}

func (ctr Controller) GetTerdaftarByUUID(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	mtd, qErr := ctr.service.GetMonitoringTerdaftar(ctx, uriParam.ID)
	if qErr != nil {
		helper.RespBadRequest(ctx, qErr.Error())
		return
	}

	helper.RespStatusOk(ctx, mtd)
}

func (ctr Controller) GetData(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
	}

	md, err := ctr.service.GetMonitoringData(ctx, uriParam.ID)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespStatusOk(ctx, md)
}

func (ctr Controller) GetAnalisa(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	id, err := uuid.Parse(uriParam.ID)
	if err != nil {
		helper.RespBadRequest(ctx, ErrQParam.Error())
		return
	}

	rowAnalisa, err := ctr.service.GetAnalisa(ctx, id)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespStatusOk(ctx, rowAnalisa)
}

func (ctr Controller) GetExport(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	id, err := uuid.Parse(uriParam.ID)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	filename, err := ctr.service.ExtractToCSV(ctx, id)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		helper.RespInternalErr(ctx, err.Error())
		return
	}

	defer file.Close()
	io.Copy(file, ctx.Request.Body)
}
