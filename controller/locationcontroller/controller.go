package locationcontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

type ILocationController interface {
	CreateProvinsi(ctx *gin.Context)
	CreateKabupaten(ctx *gin.Context)
	CreateKecamatan(ctx *gin.Context)
	CreateDesa(ctx *gin.Context)
	DeleteLocation(ctx *gin.Context)
}

type Controller struct {
	service svc.IServices
}

func New(service svc.IServices) Controller {
	return Controller{service: service}
}
