package frontend

import (
	"errors"
	"io"
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	svc "mopoen-remake/service"
	"mopoen-remake/service/servicemodel"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewMonitoringController(service svc.IServices) MonitoringController {
	return MonitoringController{service: service}
}

type MonitoringController struct {
	service svc.IServices
}

var ErrQParam = errors.New("query param tidak valid")

// GetTerdaftar is a function to find all monitoring terdaftar and filter by lokasi_id and sensor_id
// @lokasi_id is a query param to filter lokasi_id
// @sensor_id is a query param to filter sensor_id
func (ctr MonitoringController) GetTerdaftar(ctx *gin.Context) {
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

// find monitoring terdaftar by uuid
func (ctr MonitoringController) GetTerdaftarByUUID(ctx *gin.Context) {
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

func (ctr MonitoringController) GetData(ctx *gin.Context) {
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

func (ctr MonitoringController) GetAnalisa(ctx *gin.Context) {
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

// ExportAndDownload is a function to export monitoring data to csv and download it
func (ctr MonitoringController) ExportAndDownload(ctx *gin.Context) {
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

func (ctr MonitoringController) DaftarMonitoring(ctx *gin.Context) {
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

func (ctr MonitoringController) CreateValue(ctx *gin.Context) {
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
