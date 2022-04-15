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

	server = Server{
		env:    env,
		Engine: engine,
	}
	return
}

func (svr Server) Run() {
	svr.Engine.Run(svr.env.ServerAddress)
}
