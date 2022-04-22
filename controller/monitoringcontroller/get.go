package monitoringcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetTerdaftar(ctx *gin.Context) {
	Q, ok := ctx.GetQuery("lokasi")
	queryParam, convErr := strconv.Atoi(Q)
	if !ok || convErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": convErr,
		})
		return
	}

	mtd, err := ctr.service.GetMonitoringTerdaftar(ctx, int32(queryParam))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": mtd,
	})
}
