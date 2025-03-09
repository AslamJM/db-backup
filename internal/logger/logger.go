package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/AslamJM/db-backup/internal/utils"
)

const logDir = "logs"
const errorLogFile = "errors.log"

var ErrorLog *log.Logger

func GetLogger(dbname string) (*log.Logger, error) {
	var Log *log.Logger
	err := utils.EnsureDir(logDir)

	if err != nil {
		return nil, err
	}

	path := filepath.Join(logDir, fmt.Sprintf("%s.log", dbname))

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	Log = log.New(file, "", log.Ldate|log.Ltime)
	return Log, nil
}

func InitErrorLog() {
	err := utils.EnsureDir(logDir)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filepath.Join(logDir, errorLogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	ErrorLog = log.New(file, "", log.Ldate|log.Ltime)

}
