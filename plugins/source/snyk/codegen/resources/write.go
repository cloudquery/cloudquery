package resources

import (
	"fmt"
	"go/format"
	"os"
)

// write will format data as Go code & write it fo file
func write(path string, data []byte) error {
	formatted, err := format.Source(data)
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", path, err)
	} else {
		data = formatted
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}
