package locationcontroller

import (
	"errors"
	"mopoen-remake/controller/request"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetAllLokasi(ctx *gin.Context) {
	uriParam := request.UriParamTipeLokasi{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	Q, ok := ctx.GetQuery("depends")
	queryParam, err := strconv.Atoi(Q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	var lokasi interface{}

	switch uriParam.Tipe {
	case Provinsi:
		if !ok {
			lokasi, err = ctr.service.GetAllProvinsi(ctx)
		}

	case Kabupaten:
		if !ok {
			lokasi, err = ctr.service.GetAllKabupaten(ctx)
		}

		lokasi, err = ctr.service.GetLokasiBy(ctx, uriParam.Tipe, int32(queryParam))

	case Kecamatan:
		if !ok {
			lokasi, err = ctr.service.GetAllKecamatan(ctx)
		}

		lokasi, err = ctr.service.GetLokasiBy(ctx, uriParam.Tipe, int32(queryParam))

	case Desa:
		if !ok {
			lokasi, err = ctr.service.GetAllDesa(ctx)
		}

		lokasi, err = ctr.service.GetLokasiBy(ctx, uriParam.Tipe, int32(queryParam))

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
		"data": lokasi,
	})
}
