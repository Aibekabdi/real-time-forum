package main

import (
	"context"
	"fmt"
	"forum/internal/http"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"forum/pkg/postgres"
	"forum/pkg/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Preparing environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	// Load configuration from config/config.json which contains details such as DB connection params
	conf, err := utils.NewConfig("./configs/config.json")
	if err != nil {
		log.Fatalf("failed to load configs: %s", err.Error())
	}

	// Connect to the postgres DB
	db, err := postgres.NewPostgresDB(conf.Database)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Close connection postgres DB
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal("can't close connection db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()
	// preparing handler <- -> service  <- -> repository
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := http.NewHandler(service)
	// Running Server
	srv := new(server.Server)
	go func() {
		if err := srv.Run(conf.Api.Port, handler.InitRoutes()); err != nil {
			log.Printf("error occured while running http server: %s", err.Error())
			return
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	if err = srv.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
