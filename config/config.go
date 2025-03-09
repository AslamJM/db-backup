package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/AslamJM/db-backup/internal/logger"
	"github.com/AslamJM/db-backup/internal/utils"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"database"`
	OutDir   string `json:"output_dir"`
}

const confFilesDir = "conf_files"

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

// get all files names ends with .json
// get the files in the confFileDir
func GetAllConfigFiles() []string {
	err := utils.EnsureDir(confFilesDir)

	if err != nil {
		logger.ErrorLog.Println("conf_file dir: ", err)
		os.Exit(1)
	}

	out := []string{}

	files, err := os.ReadDir(confFilesDir)

	if err != nil {
		logger.ErrorLog.Println("conf_file dir read: ", err)
		os.Exit(1)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			out = append(out, filepath.Join(confFilesDir, f.Name()))
		}
	}

	return out
}
