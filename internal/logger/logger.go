package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AslamJM/db-backup/internal/utils"
)

const errorLogFile = "errors.log"

var ErrorLog *log.Logger

func getLogsDir() string {
	return os.Getenv("LOGS_DIR")
}

func GetLogger(dbname string) (*log.Logger, *os.File, error) {
	var Log *log.Logger

	logDir := getLogsDir()

	err := utils.EnsureDir(logDir)

	if err != nil {
		return nil, nil, err
	}

	path := filepath.Join(logDir, fmt.Sprintf("%s.log", dbname))

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, nil, err
	}

	Log = log.New(file, "", log.Ldate|log.Ltime)
	return Log, file, nil
}

func InitErrorLog() (*os.File, error) {

	logDir := getLogsDir()

	err := utils.EnsureDir(logDir)

	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath.Join(logDir, errorLogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	ErrorLog = log.New(file, "", log.Ldate|log.Ltime)

	return file, nil

}
