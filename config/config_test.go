package config

import (
	"os"
	"testing"
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
