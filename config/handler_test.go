package config

import (
	"github.com/Hertucktor/archive-importer/utils"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// create a temporary file with the test configuration
	configYAML := `
dbUser: "testUser"
dbPass: "testPass"
dbPort: "27017"
dbName: "testDB"
dbCollectionAllcards: "testAllCards"
dbCollectionMycards: "testMyCards"
dbCollectionSetimages: "testSetImages"
dbCollectionSetNames: "testSetNames"
`
	tmpfile, err := os.CreateTemp("", "config_test.yml")
	if err != nil {
		t.Fatalf("error creating temporary config file: %v", err)
	}
	defer func(name string) {
		if err := os.Remove(name); err != nil {
			t.Error(err)
		}
	}(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(configYAML)); err != nil {
		t.Fatalf("error writing to temporary config file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("error closing temporary config file: %v", err)
	}

	// call GetConfig with the temporary file and a logger
	logger := utils.InitializeLogger("")
	config, err := GetConfig(tmpfile.Name(), logger.Sugar())
	if err != nil {
		t.Fatalf("error getting config: %v", err)
	}

	// check that the values are correct
	if config.DBUser != "testUser" {
		t.Errorf("expected DBUser to be \"testUser\", got %q", config.DBUser)
	}
	if config.DBPass != "testPass" {
		t.Errorf("expected DBPass to be \"testPass\", got %q", config.DBPass)
	}
	if config.DBPort != "27017" {
		t.Errorf("expected DBPort to be \"27017\", got %q", config.DBPort)
	}
	if config.DBName != "testDB" {
		t.Errorf("expected DBName to be \"testDB\", got %q", config.DBName)
	}
	if config.DBCollectionAllcards != "testAllCards" {
		t.Errorf("expected DBCollectionAllcards to be \"testAllCards\", got %q", config.DBCollectionAllcards)
	}
	if config.DBCollectionMycards != "testMyCards" {
		t.Errorf("expected DBCollectionMycards to be \"testMyCards\", got %q", config.DBCollectionMycards)
	}
	if config.DBCollectionSetimages != "testSetImages" {
		t.Errorf("expected DBCollectionSetimages to be \"testSetImages\", got %q", config.DBCollectionSetimages)
	}
	if config.DBCollectionSetNames != "testSetNames" {
		t.Errorf("expected DBCollectionSetNames to be \"testSetNames\", got %q", config.DBCollectionSetNames)
	}
}
