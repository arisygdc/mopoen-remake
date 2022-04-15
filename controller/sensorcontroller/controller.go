package sensorcontroller

import (
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

type ISensorController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type Controller struct {
	service svc.IServices
}

func New(service svc.IServices) Controller {
	return Controller{service: service}
}
