package main

import (
	"forum/pkg/postgres"
	"forum/pkg/utils"
	"log"

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

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatal("can't close connection db, err:", err)
		} else {
			log.Println("db closed")
		}
	}()
}
