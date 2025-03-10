package backup

import (
	"os"
	"testing"
)

func TestBackupPostgres(t *testing.T) {
	file, err := os.CreateTemp("conf_files", "pg-test.json")

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(file.Name())

	// already existing db in local instance
	// change accordingly
	sample :=
		` 
			{
			"host": "localhost",
			"type": "pg",
			"port": 5432,
			"user": "dev",
			"password": "123456",
			"database": "feedbacksdb",
			"output_dir": "out"
			}
		`
	if err := os.WriteFile(file.Name(), []byte(sample), 0644); err != nil {
		t.Fatal(err)
	}

	BackupPostgres(file.Name())

}
