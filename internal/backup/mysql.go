package backup

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/AslamJM/db-backup/config"
)

// format the backup file name and date like this: dbname_2023-04-15.sql
func backupFileName(dbName string, ext string) string {
	date := time.Now().Format("15-04-2023")
	return dbName + "_" + date + ext
}

func BackupMySQL(cfgFileName string) error {
	cfg, err := config.GetConfigFromJSON(cfgFileName)

	if err != nil {
		return err
	}

	bfile := backupFileName(cfg.DBName, ".sql")

	cmd := exec.Command("mysqldump", "-h", cfg.Host, "-u", cfg.User, fmt.Sprintf("-p%s", cfg.Password), cfg.DBName)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	EnsureDir(fmt.Sprintf("%s/%s", outputDir, cfg.OutDir))

	filePath := fmt.Sprintf("%s/%s/%s", outputDir, cfg.OutDir, bfile)

	err = SaveToLocal(filePath, output)

	return err

}
