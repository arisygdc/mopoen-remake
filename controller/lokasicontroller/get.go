package lokasicontroller

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
	var lokasi interface{}
	var err error
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if uriParam.Tipe == Provinsi {
		lokasi, err = ctr.service.GetAllProvinsi(ctx)
	} else {

		Q, ok := ctx.GetQuery("depends")
		if ok {
			queryParam, convErr := strconv.Atoi(Q)
			if convErr != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			uriParam.Tipe = strings.ToLower(uriParam.Tipe)
			lokasi, err = ctr.service.GetLokasiBy(ctx, uriParam.Tipe, int32(queryParam))
		} else {
			switch uriParam.Tipe {
			case Kabupaten:
				lokasi, err = ctr.service.GetAllKabupaten(ctx)

			case Kecamatan:
				lokasi, err = ctr.service.GetAllKecamatan(ctx)

			case Desa:
				lokasi, err = ctr.service.GetAllDesa(ctx)

			default:
				err = errors.New("tipe lokasi tidak tersedia")
			}
		}
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": lokasi,
	})
}
