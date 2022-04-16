package main

import (
	"log"
	"mopoen-remake/config"
	"mopoen-remake/controller"
	"mopoen-remake/router"
	"mopoen-remake/server"
	svc "mopoen-remake/service"
)

func main() {
	env, err := config.New(".")
	if err != nil {
		log.Fatal(err)
	}

	service, err := svc.New(env)
	if err != nil {
		log.Fatal()
	}

	s := server.New(env)
	router.New(s.Engine, controller.New(service))
	s.Run()
}
