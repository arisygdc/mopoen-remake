package locationcontroller

import (
	"errors"
	"mopoen-remake/controller/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) CreateLocation(ctx *gin.Context) {
	uriParam := request.UriParamTipeLocation{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	var err error

	switch uriParam.Tipe {
	case ctr.provinsi:
		req := request.PostProvinsi{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err = ctr.service.CreateProvinsi(ctx, req.Nama)

	case ctr.kabupaten:
		req := request.PostKabupaten{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err = ctr.service.CreateKabupaten(ctx, req.Provinsi_id, req.Nama)

	case ctr.kecamatan:
		req := request.PostKecamatan{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err = ctr.service.CreateKecamatan(ctx, req.Kabupaten_id, req.Nama)

	case ctr.desa:
		req := request.PostDesa{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err = ctr.service.CreateDesa(ctx, req.Kecamatan_id, req.Nama)

	default:
		err = errors.New("tipe lokasi tidak tersedia")
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": uriParam.Tipe + " created",
	})
}
