package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestSimpleGoception runs a simple test to verify Goception works
func TestSimpleGoception(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "goception_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a simple test file
	code := `
		const x: int = 42;
		print(x);
	`
	testFile := filepath.Join(tempDir, "simple.gct")
	if err := os.WriteFile(testFile, []byte(code), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Run the test
	cmd := exec.Command("go", "run", "../main.go", testFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run Goception: %v\nOutput: %s", err, output)
	}

	// Check result
	outputStr := strings.TrimSpace(string(output))
	outputStr = strings.TrimSuffix(outputStr, "null") // Remove trailing null
	outputStr = strings.TrimSpace(outputStr)

	expected := "42"
	if outputStr != expected {
		t.Errorf("Expected output: '%s', got: '%s'", expected, outputStr)
	}
}
