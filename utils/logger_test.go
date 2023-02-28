package utils

import (
	"os"
	"strings"
	"testing"
)

func TestInitializeLogger(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("failed to create temporary file: %s", err)
	}
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			t.Error(err)
		}
	}(tmpfile.Name())

	// Call the function with the temporary file
	logger := InitializeLogger(tmpfile.Name())

	// Test that logger is not nil
	if logger == nil {
		t.Errorf("expected logger to not be nil")
	}

	// Test that logging works
	logger.Info("testing logger")

	// Verify that log was written to file
	b, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("failed to read log file: %s", err)
	}
	if len(b) == 0 {
		t.Errorf("expected log file to not be empty")
	}
}

func TestInitializeSugarLogger(t *testing.T) {
	// Create a temporary log file for testing
	tmpfile, err := os.CreateTemp("", "test.log")
	if err != nil {
		t.Fatalf("failed to create temporary log file: %v", err)
	}
	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			t.Error(err)
		}
	}(tmpfile.Name())

	// Initialize the logger with the temporary log file
	logger := InitializeSugarLogger(tmpfile.Name())

	// Test the logger by logging a message
	logger.Infow("Test message", "key", "value")

	// Read the contents of the log file
	b, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	// Verify that the log file contains the expected message
	if !strings.Contains(string(b), "Test message") {
		t.Errorf("log file did not contain expected message")
	}
}
