package server

import (
	"mopoen-remake/config"
	"mopoen-remake/controller/frontend"
	"mopoen-remake/controller/sensorgateway"
	"mopoen-remake/repository"
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

func (svr Server) ExposeRoute(repo repository.Repository) error {

	router := svr.Engine
	lokasiSvc := svc.NewLokasiService(repo)
	lokasiController := frontend.NewLokasiController(lokasiSvc)
	sensorSvc := svc.NewSensorService(repo)
	sensorController := frontend.NewSensorController(sensorSvc)
	monitoringSvc := svc.NewMonitoringService(repo)
	monitoringController := frontend.NewMonitoringController(monitoringSvc)

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

	sensorGSvc := svc.NewSensorGatewayService(repo)
	sensorG := sensorgateway.NewSensorGatewayController(sensorGSvc)
	sensorGRoute := router.Group("/api/sensor")
	sensorGRoute.POST("/value", sensorG.SaveDataFromSensor)
	return nil
}

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
