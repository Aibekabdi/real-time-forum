package utils

import (
	"encoding/json"
	"os"
)

type (
	Conf struct {
		Api      Api      `json:"api"`
		Client   Client   `json:"client"`
		Database Database `json:"database"`
	}
	Api struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Client struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	}
)

// loading config
func NewConfig(path string) (*Conf, error) {
	var newConfig Conf
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(file).Decode(&newConfig); err != nil {
		return nil, err
	}
	return &newConfig, nil
}
