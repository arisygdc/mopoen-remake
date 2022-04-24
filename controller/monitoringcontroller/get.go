package monitoringcontroller

import (
	"errors"
	"log"
	"mopoen-remake/controller/request"
	"net/http"
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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": qErr.Error(),
			})
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
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": qErr.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": mtd,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err,
	})

}

func (ctr Controller) GetTerdaftarByUUID(ctx *gin.Context) {
	var uriParam request.GetUUID
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	mtd, qErr := ctr.service.GetMonitoringTerdaftar(ctx, uriParam.ID)
	log.Println(mtd)
	if qErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": qErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": mtd,
	})
}

func (ctr Controller) GetData(ctx *gin.Context) {
	Q, ok := ctx.GetQuery("uuid")
	if !ok || len(Q) != 36 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrQParam,
		})
		return
	}

	md, err := ctr.service.GetMonitoringData(ctx, Q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": md,
	})
}

func (ctr Controller) GetAnalisa(ctx *gin.Context) {
	Q, ok := ctx.GetQuery("uuid")
	if !ok || len(Q) != 36 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrQParam,
		})
		return
	}

	id, err := uuid.Parse(Q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": ErrQParam,
		})
		return
	}
	rowAnalisa, err := ctr.service.GetAnalisa(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": rowAnalisa,
	})
}
