package locationcontroller

import (
	"mopoen-remake/controller/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctr Controller) CreateProvinsi(ctx *gin.Context) {
	req := request.PostProvinsi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.CreateProvinsi(ctx, req.Nama); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func (ctr Controller) CreateKabupaten(ctx *gin.Context) {
	req := request.PostKabupaten{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.CreateKabupaten(ctx, req.Provinsi_id, req.Nama); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func (ctr Controller) CreateKecamatan(ctx *gin.Context) {
	req := request.PostKecamatan{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.CreateKecamatan(ctx, req.Kabupaten_id, req.Nama); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func (ctr Controller) CreateDesa(ctx *gin.Context) {
	req := request.PostDesa{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := ctr.service.CreateDesa(ctx, req.Kecamatan_id, req.Nama); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}
