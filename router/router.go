package router

import "github.com/gin-gonic/gin"

func New(server *gin.Engine) {
	apiV1Route(server.Group("/api/v1"))
}
