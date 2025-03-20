package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestTypeError tests that type errors are detected properly
func TestTypeError(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "goception_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file with a type error
	code := `
		const age: int = "not a number";  // Should error: type mismatch
		print("If you see this, type checking failed");
	`
	testFile := filepath.Join(tempDir, "type-error.gct")
	if err := os.WriteFile(testFile, []byte(code), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Run the test
	cmd := exec.Command("go", "run", "../main.go", testFile)
	output, _ := cmd.CombinedOutput() // Ignore the error - we expect one
	outputStr := string(output)

	// Check if the error message contains "type mismatch"
	if !strings.Contains(outputStr, "type mismatch") {
		t.Errorf("Expected type mismatch error, but got: %s", outputStr)
	}
}
