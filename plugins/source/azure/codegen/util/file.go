package util

import (
	"fmt"
	"os"
	"os/exec"
)

func WriteAndFormat(filePath string, data []byte) error {
	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return FormatFile(filePath)
}

func FormatFile(filePath string) error {
	if err := goimportsFile(filePath); err != nil {
		return err
	}
	return gofumptFile(filePath)
}

func goimportsFile(filePath string) error {
	cmd := exec.Command("goimports", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func gofumptFile(filePath string) error {
	cmd := exec.Command("gofumpt", "-w", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
