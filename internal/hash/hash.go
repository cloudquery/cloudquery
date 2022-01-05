package hash

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
