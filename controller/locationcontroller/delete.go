package locationcontroller

import (
	"errors"
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) DeleteLocation(ctx *gin.Context) {
	queryParam := request.DeleteLocationType{}
	if err := ctx.ShouldBindQuery(&queryParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	req := request.DeleteLocation{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var err error
	switch queryParam.Type {
	case "provinsi":
		err = ctr.service.DeleteProvinsi(ctx, req.Id)
	case "kabupaten":
		err = ctr.service.DeleteKabupaten(ctx, req.Id)
	case "kecamatan":
		err = ctr.service.DeleteKecamatan(ctx, req.Id)
	case "desa":
		err = ctr.service.DeleteDesa(ctx, req.Id)
	default:
		err = errors.New("tipe lokasi tidak tersedia")
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": queryParam.Type + " berhasil dibuat",
	})
}
