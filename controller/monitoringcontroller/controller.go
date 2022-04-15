package monitoringcontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

type IMonitoringController interface {
	CreateDaftar(ctx *gin.Context)
	CreateValue(ctx *gin.Context)
}

type Controller struct {
	service svc.IServices
}

func New(service svc.IServices) Controller {
	return Controller{service: service}
}
