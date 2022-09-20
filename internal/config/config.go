package config

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
		Port string `json:"port"`
	}
	Client struct {
		Port string `json:"port"`
	}
	Database struct {
		Driver   string `json:"driver"`
		Path     string `json:"path"`
		FileName string `json:"fileName"`
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
	return &config, nil
}
