package server

import (
	"mopoen-remake/config"
	"mopoen-remake/server/middleware"

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

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
