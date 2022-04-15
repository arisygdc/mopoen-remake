package router

import (
	"mopoen-remake/controller"

	"github.com/gin-gonic/gin"
)

func New(server *gin.Engine, controller controller.Controller) {
	apiV1Route(server.Group("/api/v1"), controller)
}
