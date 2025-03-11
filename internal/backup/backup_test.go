package backup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBackupPostgres(t *testing.T) {
	file, err := os.Create(filepath.Join("conf_files", "pg-test.json"))

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
			"user": "db_admin",
			"password": "123456",
			"database": "leave_app",
			"output_dir": "out"
			}
		`
	if err := os.WriteFile(file.Name(), []byte(sample), 0644); err != nil {
		t.Fatal(err)
	}

	BackupPostgres(file.Name())

	expectedOutput := filepath.Join("outputs", "leave_app")

	dir, err := os.ReadDir(expectedOutput)
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(expectedOutput)

	if len(dir) == 0 {
		t.Error("expected directory to be not empty")
	}

}
