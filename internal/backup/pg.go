package backup

import (
	"fmt"
	"os/exec"

	"github.com/AslamJM/db-backup/config"
	"github.com/AslamJM/db-backup/internal/logger"
)

func BackupPostgres(cfgFileName string) {
	cfg, err := config.GetConfigFromJSON(cfgFileName)

	if err != nil {
		logger.ErrorLog.Printf("error reading config %s : %v\n", cfgFileName, err)
		return
	}

	dbLog, file, err := logger.GetLogger(cfg.DBName)

	if err != nil {
		logger.ErrorLog.Printf("error getting logger for %s : %v\n", cfg.DBName, err)
		return
	}

	defer file.Close()

	bfile := backupFileName(cfg.DBName, ".sql")

	cmd := exec.Command("pg_dump", "-h", cfg.Host, "-p", fmt.Sprintf("%d", cfg.Port), "-U", cfg.User, "-F", "c", "-b", "-v", "-f", bfile, cfg.DBName)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", cfg.Password))

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
