package main

import (
	"log"
	"os"

	"github.com/AslamJM/db-backup/internal/backup"
	"github.com/AslamJM/db-backup/internal/logger"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := logger.InitErrorLog()

	if err != nil {
		log.Fatal(err)
	}

	c := cron.New()
	_, err = c.AddFunc(os.Getenv("CRON_TIME"), func() {
		backup.RunConcurrentBackups()
	})

	if err != nil {
		logger.ErrorLog.Printf("❌ failed to schedule backup: %v", err)
		os.Exit(1)
	}

	c.Start()

	// close the error log file
	defer file.Close()

	select {}

}
