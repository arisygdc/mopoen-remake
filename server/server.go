package server

import (
	"mopoen-remake/config"
	"mopoen-remake/controller/frontend"
	"mopoen-remake/controller/sensorgateway"
	"mopoen-remake/pkg/mail"
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

	// Mail library
	mailSvc := mail.NewMailSender(svr.env.GmailUser, svr.env.GmailPass)

	lokasiSvc := svc.NewLokasiService(repo)
	lokasiController := frontend.NewLokasiController(lokasiSvc)
	sensorSvc := svc.NewSensorService(repo)
	sensorController := frontend.NewSensorController(sensorSvc)
	monitoringSvc := svc.NewMonitoringService(repo, mailSvc)
	monitoringController := frontend.NewMonitoringController(monitoringSvc)

	ruoterApiV1 := router.Group("/api/v1")
	// request PostSensor
	ruoterApiV1.POST("/sensor", sensorController.CreateNewTipeSensor)
	// request DeleteSensor
	ruoterApiV1.DELETE("/sensor", sensorController.DeleteTipeSensor)
	// request GetSensor
	ruoterApiV1.GET("/sensor/:id", sensorController.GetTipeSensorByID)
	// no request
	ruoterApiV1.GET("/sensors", sensorController.GetAllTipeSensor)
	// request PostNamaLokasi
	ruoterApiV1.POST("/lokasi/provinsi", lokasiController.CreateLokasiProvinsi)
	// request UriParamLokasiDepends, PostNamaLokasi
	ruoterApiV1.POST("/lokasi/:tipe/:depends", lokasiController.CreateLokasiDepends)
	// request UriParamTipeLokasi, DeleteLokasi
	ruoterApiV1.DELETE("/lokasi/:tipe", lokasiController.DeleteLokasi)
	// request UriParamTipeLokasi
	ruoterApiV1.GET("/lokasi/:tipe", lokasiController.GetAllLokasiWithType)

	ruoterApiV1.GET("/lokasi/parent", lokasiController.GetParentLokasi)
	// request PostDaftarMonitoring
	ruoterApiV1.POST("/monitoring/daftar", monitoringController.DaftarMonitoring)
	// request query lokasi_id, sensor_id
	ruoterApiV1.GET("/monitoring/terdaftar", monitoringController.GetTerdaftar)
	// request GetUUID
	ruoterApiV1.GET("/monitoring/terdaftar/:uuid", monitoringController.GetTerdaftarByUUID)
	// request GetUUID
	ruoterApiV1.GET("/monitoring/value/:uuid", monitoringController.GetData)
	// request GetUUID
	ruoterApiV1.GET("/monitoring/analisa/:uuid", monitoringController.GetAnalisa)

	// request GetUUID
	// download
	router.GET("/download/csv/:uuid", monitoringController.ExportAndDownload)

	sensorGSvc := svc.NewSensorGatewayService(repo)
	sensorG := sensorgateway.NewSensorGatewayController(sensorGSvc)
	sensorGRoute := router.Group("/api/sensor")
	sensorGRoute.POST("/value", sensorG.SaveDataFromSensor)
	return nil
}

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
