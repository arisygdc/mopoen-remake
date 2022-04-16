package router

import (
	"mopoen-remake/controller"

	"github.com/gin-gonic/gin"
)

func apiV1Route(route *gin.RouterGroup, controller controller.Controller) {
	route.POST("/sensor", controller.Sensor.Create)
	route.DELETE("/sensor", controller.Sensor.Delete)
	route.GET("/sensor", controller.Sensor.Get)
	route.GET("/sensors", controller.Sensor.GetAll)
	route.POST("/lokasi/:tipe", controller.Location.CreateLocation)
	route.DELETE("/lokasi/:tipe", controller.Location.DeleteLocation)
	route.POST("/monitoring/daftar", controller.Monitoring.CreateDaftar)
	route.POST("/monitoring/value", controller.Monitoring.CreateValue)
}
