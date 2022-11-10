package main

import (
	"forum/internal/config"
	"forum/internal/http"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// preparation configs
	conf, err := config.NewConfig("./configs/config.json")
	if err != nil {
		log.Fatalf("error occured while trying to parse config: %s", err.Error())
	}
	// preparation db
	db, err := repository.NewPostgresDB(conf.Database)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := http.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(conf.Api.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
