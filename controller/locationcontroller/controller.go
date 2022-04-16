package locationcontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

type ILocationController interface {
	CreateLocation(ctx *gin.Context)
	DeleteLocation(ctx *gin.Context)
}

type Controller struct {
	service   svc.IServices
	provinsi  string
	kabupaten string
	kecamatan string
	desa      string
}

func New(service svc.IServices) Controller {
	return Controller{
		service:   service,
		provinsi:  "provinsi",
		kabupaten: "kabupaten",
		kecamatan: "kecamatan",
		desa:      "desa",
	}
}
