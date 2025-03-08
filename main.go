package main

import (
	"fmt"
	"os"

	"github.com/AslamJM/db-backup/internal/backup"
)

func main() {
	err := backup.BackupMySQL("config.json")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Backup completed successfully")

}
