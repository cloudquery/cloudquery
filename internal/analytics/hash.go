package analytics

import (
	"crypto/sha256"
	"fmt"
)

// HashAttribute creates a one-way hash from an attribute
func HashAttribute(value string) string {
	s := sha256.New()
	_, _ = s.Write([]byte(value))
	return fmt.Sprintf("%0x", s.Sum(nil))
}
