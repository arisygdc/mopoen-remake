package monitoringcontroller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrTerdaftarQParam = errors.New("query param tidak valid")

func (ctr Controller) GetTerdaftar(ctx *gin.Context) {
	QLok, okLok := ctx.GetQuery("lokasi")
	QUUID, okUUID := ctx.GetQuery("uuid")
	queryParam, err := strconv.Atoi(QLok)

	if okLok && okUUID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": ErrTerdaftarQParam,
		})
	}

	if okLok && err == nil {
		mtd, qErr := ctr.service.GetMonitoringTerdaftarByLokasi(ctx, int32(queryParam))
		if qErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": mtd,
		})
		return
	}

	if okUUID {
		mtd, qErr := ctr.service.GetMonitoringTerdaftar(ctx, QUUID)
		if qErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": mtd,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": ErrTerdaftarQParam,
	})

}

func (ctr Controller) GetData(ctx *gin.Context) {
	Q, ok := ctx.GetQuery("uuid")
	if !ok || len(Q) != 36 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "uuid tidak valid",
		})
		return
	}

	md, err := ctr.service.GetMonitoringData(ctx, Q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": md,
	})
}
