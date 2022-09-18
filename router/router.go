package router

import (
	"mopoen-remake/controller"
	"mopoen-remake/server/middleware"

	"github.com/gin-gonic/gin"
)

func New(server gin.IRouter, controller controller.Controller) {
	ruoterApiV1 := server.Group("/api/v1").Use(middleware.Bearear())
	apiV1Route(ruoterApiV1, controller)
	apiSensor(server.Group("/api/sensor"), controller.Monitoring)
}
