package server

import (
	"mopoen-remake/config"
	"mopoen-remake/controller/frontend"
	"mopoen-remake/controller/sensorgateway"
	"mopoen-remake/server/middleware"
	svc "mopoen-remake/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env    config.Environment
	Engine *gin.Engine
}

func New(env config.Environment) (server Server) {
	gin.SetMode(env.ServerEnv)
	engine := gin.Default()
	engine.Use(middleware.Cors())

	server = Server{
		env:    env,
		Engine: engine,
	}

	return
}

func (svr Server) ExposeRoute() error {
	router := svr.Engine
	service, err := svc.New(svr.env)
	if err != nil {
		return err
	}

	lokasiController := frontend.NewLokasiController(service)
	sensorController := frontend.NewSensorController(service)
	monitoringController := frontend.NewMonitoringController(service)

	ruoterApiV1 := router.Group("/api/v1").Use(middleware.Bearear())
	ruoterApiV1.POST("/sensor", sensorController.CreateNewTipeSensor)
	ruoterApiV1.DELETE("/sensor", sensorController.DeleteTipeSensor)
	ruoterApiV1.GET("/sensor/:id", sensorController.GetTipeSensorByID)
	ruoterApiV1.GET("/sensors", sensorController.GetAllTipeSensor)
	ruoterApiV1.POST("/lokasi/:tipe", lokasiController.CreateLokasi)
	ruoterApiV1.DELETE("/lokasi/:tipe", lokasiController.DeleteLokasi)
	ruoterApiV1.GET("/lokasi/:tipe", lokasiController.GetAllLokasiWithType)
	ruoterApiV1.POST("/monitoring/daftar", monitoringController.DaftarMonitoring)
	ruoterApiV1.GET("/monitoring/terdaftar", monitoringController.GetTerdaftar)
	ruoterApiV1.GET("/monitoring/terdaftar/:uuid", monitoringController.GetTerdaftarByUUID)
	ruoterApiV1.GET("/monitoring/value/:uuid", monitoringController.GetData)
	ruoterApiV1.GET("/monitoring/analisa/:uuid", monitoringController.GetAnalisa)
	ruoterApiV1.GET("/monitoring/csv/:uuid", monitoringController.ExportAndDownload)

	sensorG := sensorgateway.New(service)
	sensorGRoute := router.Group("/api/sensor")
	sensorGRoute.POST("/value", sensorG.SaveDataFromSensor)
	return nil
}

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
