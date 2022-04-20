package locationcontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

var Provinsi = "provinsi"
var Kabupaten = "kabupaten"
var Kecamatan = "kecamatan"
var Desa = "desa"

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
