package controller

import (
	"mopoen-remake/controller/locationcontroller"
	"mopoen-remake/controller/monitoringcontroller"
	"mopoen-remake/controller/sensorcontroller"
	svc "mopoen-remake/service"
)

type Controller struct {
	Sensor     sensorcontroller.ISensorController
	Location   locationcontroller.ILocationController
	Monitoring monitoringcontroller.IMonitoringController
}

func New(service svc.IServices) Controller {
	return Controller{
		Sensor:     sensorcontroller.New(service),
		Location:   locationcontroller.New(service),
		Monitoring: monitoringcontroller.New(service),
	}
}
