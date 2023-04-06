package main

import (
	"log"
	"mopoen-remake/config"
	"mopoen-remake/repository"
	"mopoen-remake/server"
)

func main() {
	env, err := config.New(".")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.NewRepository(env.DBDriver, env.DBSource)

	s := server.New(env)
	s.ExposeRoute(repo)
	s.Run()
}
