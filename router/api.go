package router

import (
	"mopoen-remake/controller"

	"github.com/gin-gonic/gin"
)

func apiV1Route(route gin.IRouter, controller controller.Controller) {
	route.POST("/sensor", controller.Sensor.Create)
	route.DELETE("/sensor", controller.Sensor.Delete)
	route.GET("/sensor/:id", controller.Sensor.Get)
	route.GET("/sensors", controller.Sensor.GetAll)
	route.POST("/lokasi/:tipe", controller.Location.CreateLokasi)
	route.DELETE("/lokasi/:tipe", controller.Location.DeleteLokasi)
	route.GET("/lokasi/:tipe", controller.Location.GetAllLokasi)
	route.POST("/monitoring/daftar", controller.Monitoring.CreateDaftar)
	route.POST("/monitoring/value", controller.Monitoring.CreateValue)
	route.GET("/monitoring/terdaftar", controller.Monitoring.GetTerdaftar)
	route.GET("/monitoring/terdaftar/:uuid", controller.Monitoring.GetTerdaftarByUUID)
	route.GET("/monitoring/value/:uuid", controller.Monitoring.GetData)
	route.GET("/monitoring/analisa/:uuid", controller.Monitoring.GetAnalisa)
}
