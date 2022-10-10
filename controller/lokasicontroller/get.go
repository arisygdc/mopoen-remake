package lokasicontroller

import (
	"database/sql"
	"errors"
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetAllLokasi(ctx *gin.Context) {
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
