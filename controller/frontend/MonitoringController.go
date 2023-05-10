package frontend

import (
	"errors"
	"io"
	"mopoen-remake/controller/helper"
	"mopoen-remake/request"
	ifM "mopoen-remake/service/serviceInterface"
	"mopoen-remake/service/servicemodel"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewMonitoringController(service ifM.MonitoringInterface) MonitoringController {
	return MonitoringController{service: service}
}

type MonitoringController struct {
	service ifM.MonitoringInterface
}

var ErrQParam = errors.New("query param tidak valid")

// GetTerdaftar is a function to find all monitoring terdaftar
// query @lokasi_id is to filter lokasi_id
// query @sensor_id is to filter sensor_id
func (ctr MonitoringController) GetTerdaftar(ctx *gin.Context) {
	QLok, okLok := ctx.GetQuery("lokasi_id")
	QSensor, okSensor := ctx.GetQuery("sensor_id")
	lok_id, _ := strconv.Atoi(QLok)
	sensor_id, _ := strconv.Atoi(QSensor)
	var err error
	if okLok || okSensor {
		mtd, err := ctr.service.GetMonitoringTerdaftar(ctx, &servicemodel.GetMonitoringTerdaftarFilterOptions{
			LokasiID:     int32(lok_id),
			TipeSensorID: int32(sensor_id),
		})
		if err != nil {
			helper.RespCatchSqlErr(ctx, err)
			return
		}
		helper.RespStatusOk(ctx, mtd)
		return
	}

	mtd, err := ctr.service.GetMonitoringTerdaftar(ctx, nil)
	if err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}
	helper.RespStatusOk(ctx, mtd)
}

// find monitoring terdaftar by uuid
func (ctr MonitoringController) GetTerdaftarByUUID(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	id, err := uuid.Parse(uriParam.ID)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
	}

	mtd, err := ctr.service.GetMonitoringTerdaftarByID(ctx, id)
	if err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	ctx.JSON(200, mtd)
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
	ctx.JSON(200, rowAnalisa)
}

// ExportAndDownload is a function to export monitoring data to csv and download it
func (ctr MonitoringController) ExportAndDownload(ctx *gin.Context) {
	var uriParam request.GetFile
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}
	format := ".csv"

	if !strings.HasSuffix(uriParam.FileName, format) {
		helper.RespBadRequest(ctx, "format tidak didukung")
	}

	id, err := uuid.Parse(uriParam.FileName[:len(uriParam.FileName)-len(format)])
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	filebuffer, err := ctr.service.GetCsvBuffer(ctx, id)
	if err != nil {
		helper.RespInternalErr(ctx, err.Error())
		return
	}

	data, err := io.ReadAll(filebuffer)
	if err != nil {
		helper.RespInternalErr(ctx, err.Error())
		return
	}

	filename := uriParam.FileName

	ctx.Data(200, "text/csv", data)
	ctx.Request.Header.Add("Content-Disposition", "attachment; filename="+filename)
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

	helper.RespCreated(ctx, req.Nama+" created")
}
