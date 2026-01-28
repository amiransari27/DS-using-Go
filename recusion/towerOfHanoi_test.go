// filepath: /Users/amir/Workspace/DS-using-Go/recusion/towerOfHanoi_test.go
package recusion

import (
	"bytes"
	"os"
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	originalStdout := os.Stdout                   // Save the original stdout
	defer func() { os.Stdout = originalStdout }() // Restore original stdout after test

	// Create a temporary file to capture output
	tempFile, err := os.CreateTemp("", "towerOfHanoi_output")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	os.Stdout = tempFile // Redirect output to temp file

	n := 3
	expectedOutput := "Move disk 1 from rod A to rod C\nMove disk 2 from rod A to rod B\nMove disk 1 from rod C to rod B\nMove disk 3 from rod A to rod C\nMove disk 1 from rod B to rod A\nMove disk 2 from rod B to rod C\nMove disk 1 from rod A to rod C\n"

	count := TowerOfHanoi(n, "A", "C", "B")

	if count != 7 {
		t.Errorf("Expected 7 moves, got %d", count)
	}

	// Read the output from the temp file
	tempFile.Seek(0, 0) // Reset file pointer to the beginning
	output := new(bytes.Buffer)
	output.ReadFrom(tempFile)

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
