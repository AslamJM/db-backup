package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/AslamJM/db-backup/internal/utils"
)

func TestGetConfigFromJSON(t *testing.T) {
	file, err := os.CreateTemp("", "config.json")

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(file.Name())

	sample :=
		`{
				"host": "localhost",
				"port": 3306,
				"user":"root",
				"password":"root",
				"database":"test"
		 }	
			`

	err = os.WriteFile(file.Name(), []byte(sample), 0644)

	if err != nil {
		t.Fatal(err)
	}

	config, err := GetConfigFromJSON(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	if config.Host != "localhost" {
		t.Errorf("Expected %s, got %s", "localhost", config.Host)
	}

	if config.Port != 3306 {
		t.Errorf("Expected %d, got %d", 3306, config.Port)
	}

	if config.User != "root" {
		t.Errorf("Expected %s, got %s", "root", config.User)
	}

	if config.Password != "root" {
		t.Errorf("Expected %s, got %s", "root", config.Password)
	}

	if config.DBName != "test" {
		t.Errorf("Expected %s, got %s", "test", config.DBName)
	}
}

func TestGetAllConfigFiles(t *testing.T) {
	var tempFiles = []string{
		"test1.json", "test2", "test3.json", "test4.txt",
	}

	if err := utils.EnsureDir(confFilesDir); err != nil {
		t.Fatal(err)
	}

	for _, tf := range tempFiles {
		f, err := os.Create(filepath.Join(confFilesDir, tf))
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
	}
	cfdir, _ := filepath.Abs(confFilesDir)
	defer os.RemoveAll(cfdir)

	expected := []string{
		filepath.Join(confFilesDir, "test1.json"),
		filepath.Join(confFilesDir, "test3.json"),
	}
	out := GetAllConfigFiles()

	if !reflect.DeepEqual(expected, out) {
		t.Errorf("test failed: expected: %v got: %v\n", expected, out)
	}

}
