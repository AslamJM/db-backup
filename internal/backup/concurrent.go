package backup

import (
	"sync"

	"github.com/AslamJM/db-backup/config"
	"github.com/AslamJM/db-backup/internal/logger"
)

func RunBackup(cfgName string) {
	cfg, err := config.GetConfigFromJSON(cfgName)

	if err != nil {
		logger.ErrorLog.Print("❌ error reading config %s: %v", cfgName, err)
	}

	if cfg.Type == "pg" {
		BackupPostgres(cfg)
	}

	if cfg.Type == "mysql" {
		BackupMySQL(cfg)
	}
}

func RunConcurrentBackups() {
	cfgFiles := config.GetAllConfigFiles()
	var wg sync.WaitGroup

	for _, cfg := range cfgFiles {
		wg.Add(1)
		go func(cfgN string) {
			defer wg.Done()
			RunBackup(cfgN)
		}(cfg)
	}
	wg.Wait()
}
