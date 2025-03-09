package backup

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/AslamJM/db-backup/config"
	"github.com/AslamJM/db-backup/internal/logger"
)

// format the backup file name and date like this: dbname_2023-04-15.sql
func backupFileName(dbName string, ext string) string {
	date := time.Now().Format("02-01-2006")
	return dbName + "_" + date + ext
}

func BackupMySQL(cfgFileName string) {

	errLogger := logger.ErrorLog

	cfg, err := config.GetConfigFromJSON(cfgFileName)

	if err != nil {
		errLogger.Printf("error reading config: %s: %v\n", cfgFileName, err)
	}

	dbLog, err := logger.GetLogger(cfg.DBName)

	if err != nil {
		errLogger.Printf("error getting logger: %s: %v\n", cfg.DBName, err)
	}

	bfile := backupFileName(cfg.DBName, ".sql")

	cmd := exec.Command("mysqldump", "-h", cfg.Host, "-u", cfg.User, fmt.Sprintf("-p%s", cfg.Password), cfg.DBName)

	output, err := cmd.CombinedOutput()

	if err != nil {
		dbLog.Println("backup failed: ", err)
	}

	EnsureDir(fmt.Sprintf("%s/%s", outputDir, cfg.OutDir))

	filePath := fmt.Sprintf("%s/%s/%s", outputDir, cfg.OutDir, bfile)

	err = SaveToLocal(filePath, output)

	if err != nil {
		dbLog.Println("error saving backup: ", err)
	} else {
		dbLog.Println("successfully done backup")
	}

}
