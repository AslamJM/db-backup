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

	cfg, err := config.GetConfigFromJSON(cfgFileName)

	if err != nil {
		logger.ErrorLog.Printf("error reading config %s : %v\n", cfgFileName, err)
		return
	}

	dbLog, file, err := logger.GetLogger(cfg.DBName)

	// early returns
	if err != nil {
		logger.ErrorLog.Printf("error getting logger for %s : %v\n", cfg.DBName, err)
		return
	}
	defer file.Close()

	bfile := backupFileName(cfg.DBName, ".sql")

	// append the db user password to the command env
	// otherwise it will prompt even if you add to the -p%pw
	cmd := exec.Command("mysqldump", "-h", cfg.Host, "-u", cfg.User, cfg.DBName)
	cmd.Env = append(cmd.Env, fmt.Sprintf("MYSQL_PWD=%s", cfg.Password))

	output, err := cmd.CombinedOutput()

	if err != nil {
		dbLog.Println("❌ backup failed: ", err)
	}

	EnsureDir(fmt.Sprintf("%s/%s", outputDir, cfg.OutDir))

	filePath := fmt.Sprintf("%s/%s/%s", outputDir, cfg.OutDir, bfile)

	err = SaveToLocal(filePath, output)

	if err != nil {
		dbLog.Println("❌ error saving backup: ", err)
	} else {
		dbLog.Println("✅ successfully done backup")
	}

}
