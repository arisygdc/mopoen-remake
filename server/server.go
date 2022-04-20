package server

import (
	"mopoen-remake/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	env    config.Environment
	Engine *gin.Engine
}

func New(env config.Environment) (server Server) {
	gin.SetMode(env.ServerEnv)
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	server = Server{
		env:    env,
		Engine: engine,
	}
	return
}

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
