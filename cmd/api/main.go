package main

import (
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	log.Println(conf)
}
