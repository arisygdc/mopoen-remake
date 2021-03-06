package locationcontroller

import (
	"errors"
	"mopoen-remake/controller/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) DeleteLokasi(ctx *gin.Context) {
	uriParam := request.UriParamTipeLokasi{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	req := request.DeleteLokasi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var err error
	switch strings.ToLower(uriParam.Tipe) {
	case Provinsi:
		err = ctr.service.DeleteProvinsi(ctx, req.Id)
	case Kabupaten:
		err = ctr.service.DeleteKabupaten(ctx, req.Id)
	case Kecamatan:
		err = ctr.service.DeleteKecamatan(ctx, req.Id)
	case Desa:
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
		"message": uriParam.Tipe + " berhasil dibuat",
	})
}
