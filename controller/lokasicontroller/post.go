package lokasicontroller

import (
	"errors"
	"fmt"
	"mopoen-remake/controller/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) CreateLokasi(ctx *gin.Context) {
	uriParam := request.UriParamTipeLokasi{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	var nama string
	var err error

	switch uriParam.Tipe {
	case Provinsi:
		req := request.PostProvinsi{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		nama = req.Nama
		err = ctr.service.CreateProvinsi(ctx, nama)

	case Kabupaten:
		req := request.PostKabupaten{}
		if err = ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		nama = req.Nama
		err = ctr.service.CreateKabupaten(ctx, req.Provinsi_id, nama)

	case Kecamatan:
		req := request.PostKecamatan{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		nama = req.Nama
		err = ctr.service.CreateKecamatan(ctx, req.Kabupaten_id, nama)

	case Desa:
		req := request.PostDesa{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			break
		}
		nama = req.Nama
		err = ctr.service.CreateDesa(ctx, req.Kecamatan_id, nama)

	default:
		err = errors.New("tipe lokasi tidak tersedia")
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("%s %s created", uriParam.Tipe, nama),
	})
}
