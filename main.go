package main

import (
	"log"
	"mopoen-remake/config"
	"mopoen-remake/server"
)

func main() {
	env, err := config.New(".")
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(env)
	s.ExposeRoute()
	s.Run()
}
