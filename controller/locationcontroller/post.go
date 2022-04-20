package locationcontroller

import (
	"errors"
	"mopoen-remake/controller/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) CreateLokasi(ctx *gin.Context) {
	uriParam := request.UriParamTipeLokasi{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	var err error

	switch uriParam.Tipe {
	case Provinsi:
		req := request.PostProvinsi{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		err = ctr.service.CreateProvinsi(ctx, req.Nama)

	case Kabupaten:
		req := request.PostKabupaten{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		err = ctr.service.CreateKabupaten(ctx, req.Provinsi_id, req.Nama)

	case Kecamatan:
		req := request.PostKecamatan{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		err = ctr.service.CreateKecamatan(ctx, req.Kabupaten_id, req.Nama)

	case Desa:
		req := request.PostDesa{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
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
