package frontend

import (
	"database/sql"
	"errors"
	"fmt"
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	ifL "mopoen-remake/service/serviceInterface"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	Provinsi  = "provinsi"
	Kabupaten = "kabupaten"
	Kecamatan = "kecamatan"
	Desa      = "desa"
)

func NewLokasiController(service ifL.LokasiInterface) LokasiController {
	return LokasiController{
		service: service,
	}
}

type LokasiController struct {
	service ifL.LokasiInterface
}

func (ctr LokasiController) CreateLokasi(ctx *gin.Context) {
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

func (ctr LokasiController) DeleteLokasi(ctx *gin.Context) {
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
	var nama string
	switch strings.ToLower(uriParam.Tipe) {
	case Provinsi:
		nama, err = ctr.service.DeleteProvinsi(ctx, req.Id)
	case Kabupaten:
		nama, err = ctr.service.DeleteKabupaten(ctx, req.Id)
	case Kecamatan:
		nama, err = ctr.service.DeleteKecamatan(ctx, req.Id)
	case Desa:
		nama, err = ctr.service.DeleteDesa(ctx, req.Id)
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
		"message": fmt.Sprintf("%s %s deleted", uriParam.Tipe, nama),
	})
}

// GetAllLokasi get all lokasi filtered by tipe
func (ctr LokasiController) GetAllLokasiWithType(ctx *gin.Context) {
	var uriParam request.UriParamTipeLokasi
	var lokasi interface{}
	var err error

	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	if uriParam.Tipe != Provinsi {
		Q, ok := ctx.GetQuery("depends")
		if ok {
			queryParam, convErr := strconv.Atoi(Q)
			if convErr == nil {
				uriParam.Tipe = strings.ToLower(uriParam.Tipe)
				lokasi, err = ctr.service.GetLokasiBy(ctx, uriParam.Tipe, int32(queryParam))
			}
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

	} else {
		lokasi, err = ctr.service.GetAllProvinsi(ctx)
	}

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			helper.RespNotFound(ctx, err.Error())
		case sql.ErrConnDone:
			helper.RespInternalErr(ctx, err.Error())
		default:
			helper.RespBadRequest(ctx, err.Error())
		}
		return
	}

	helper.RespStatusOk(ctx, lokasi)
}
