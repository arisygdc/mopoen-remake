package lokasicontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

const (
	Provinsi  = "provinsi"
	Kabupaten = "kabupaten"
	Kecamatan = "kecamatan"
	Desa      = "desa"
)

type ILocationController interface {
	CreateLokasi(ctx *gin.Context)
	DeleteLokasi(ctx *gin.Context)
	GetAllLokasi(ctx *gin.Context)
}

type Controller struct {
	service svc.IServices
}

func New(service svc.IServices) Controller {
	return Controller{
		service: service,
	}
}
