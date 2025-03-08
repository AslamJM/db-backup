package config

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"database"`
	OutDir   string `json:"output_dir"`
}

func GetConfigFromJSON(fileName string) (*DBConfig, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config DBConfig
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
