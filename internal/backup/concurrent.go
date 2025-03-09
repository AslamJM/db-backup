package backup

import (
	"sync"

	"github.com/AslamJM/db-backup/config"
)

func RunConcurrentBackups() {
	cfgFiles := config.GetAllConfigFiles()
	var wg sync.WaitGroup

	for _, cfg := range cfgFiles {
		wg.Add(1)
		go func(cfgN string) {
			defer wg.Done()
			BackupMySQL(cfgN)
		}(cfg)
	}
	wg.Wait()
}
