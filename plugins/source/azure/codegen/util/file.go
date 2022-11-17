package util

import (
	"fmt"
	"os"
	"os/exec"
)

func WriteAndFormat(filePath string, data []byte) error {
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return FormatFile(filePath)
}

func FormatFile(filePath string) error {
	cmd := exec.Command("goimports", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
