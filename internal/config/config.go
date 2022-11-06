package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Conf struct {
		Api      Api      `json:"api"`
		Database Database `json:"database"`
	}
	Api struct {
		Port string `json:"port"`
	}
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	}
)

func NewConfig(configPath string) (*Conf, error) {
	var config Conf
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variable: %s", err.Error())
	}
	config.Database.Password = os.Getenv("DB_PASSWORD")
	return &config, nil
}
