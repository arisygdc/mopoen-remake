package router

import (
	"github.com/gin-gonic/gin"
)

type IRouteSensor interface {
	CreateValue(ctx *gin.Context)
}

func apiSensor(route gin.IRoutes, controller IRouteSensor) {
	route.POST("/monitoring/value", controller.CreateValue)
}
