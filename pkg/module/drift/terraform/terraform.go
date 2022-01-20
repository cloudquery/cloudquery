package terraform

import (
	"encoding/json"
	"fmt"
	"io"
)

// LoadState loads the given reader into tfstate and validates the state version
func LoadState(reader io.Reader) (*Data, error) {
	var s Data
	if err := json.NewDecoder(reader).Decode(&s.State); err != nil {
		return nil, fmt.Errorf("invalid tf state file: %w", err)
	}
	if s.State.Version != StateVersion {
		return nil, fmt.Errorf("unsupported state version %d", s.State.Version)
	}
	return &s, nil
}
