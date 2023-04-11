package frontend

import (
	"fmt"
	"mopoen-remake/controller/helper"
	"mopoen-remake/request"
	ifL "mopoen-remake/service/serviceInterface"
	"mopoen-remake/service/servicemodel"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewLokasiController(service ifL.LokasiInterface) LokasiController {
	return LokasiController{
		service: service,
	}
}

type LokasiController struct {
	service ifL.LokasiInterface
}

func (ctr LokasiController) CreateLokasiDepends(ctx *gin.Context) {
	uriParam := request.UriParamLokasiDepends{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	req := request.PostNamaLokasi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	tipe := servicemodel.LokasiType(uriParam.Tipe)

	if tipe == servicemodel.LokKabupaten || tipe == servicemodel.LokKecamatan || tipe == servicemodel.LokDesa {
		err := ctr.service.CreateLokasi(ctx, tipe, req.Nama, uriParam.Depends)
		if err != nil {
			helper.RespBadRequest(ctx, err.Error())
			return
		}
		helper.RespCreated(ctx, fmt.Sprintf("%s %s created", uriParam.Tipe, req.Nama))
		return
	}

	helper.RespNotFound(ctx, "tipe lokasi tidak ditemukan")
}

func (ctr LokasiController) CreateLokasiProvinsi(ctx *gin.Context) {
	req := request.PostNamaLokasi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	err := ctr.service.CreateLokasi(ctx, servicemodel.LokProvinsi, req.Nama)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespCreated(ctx, fmt.Sprintf("%s %s created", servicemodel.LokProvinsi, req.Nama))
}

func (ctr LokasiController) DeleteLokasi(ctx *gin.Context) {
	uriParam := request.UriParamTipeLokasi{}
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	req := request.DeleteLokasi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	var err error
	var nama string
	nama, err = ctr.service.DeleteLokasi(ctx, servicemodel.LokasiType(uriParam.Tipe), req.Id)

	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespStatusOk(ctx, fmt.Sprintf("%s %s deleted", uriParam.Tipe, nama))
}

func (ctr LokasiController) GetParentLokasi(ctx *gin.Context) {
	lok, err := ctr.service.GetParentLokasi(ctx)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}
	helper.RespStatusOk(ctx, lok)
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

	uriParam.Tipe = strings.ToLower(uriParam.Tipe)
	tipe := servicemodel.LokasiType(uriParam.Tipe)
	Q, ok := ctx.GetQuery("depends")
	if tipe != servicemodel.LokProvinsi && ok {
		queryParam, convErr := strconv.Atoi(Q)
		if convErr == nil {
			helper.RespNotFound(ctx, "query depends harus berisi id yang valid")
			return
		}

		uriParam.Tipe = strings.ToLower(uriParam.Tipe)
		lokasi, err = ctr.service.GetLokasiBy(ctx, servicemodel.LokasiType(uriParam.Tipe), int32(queryParam))
		if err != nil {
			helper.RespCatchSqlErr(ctx, err)
			return
		}
		helper.RespStatusOk(ctx, lokasi)
		return
	}

	switch tipe {
	case servicemodel.LokProvinsi:
		lokasi, err = ctr.service.GetAllProvinsi(ctx)
	case servicemodel.LokKabupaten:
		lokasi, err = ctr.service.GetAllKabupaten(ctx)
	case servicemodel.LokKecamatan:
		lokasi, err = ctr.service.GetAllKecamatan(ctx)
	case servicemodel.LokDesa:
		lokasi, err = ctr.service.GetAllDesa(ctx)
	}

	if err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	helper.RespStatusOk(ctx, lokasi)
}
