package terraform

import (
	"encoding/json"
	"fmt"
	"io"
)

// parseAndValidate received reader turn in into Data state and validate the state version
func ParseAndValidate(reader io.Reader) (*Data, error) {
	var s Data
	if err := json.NewDecoder(reader).Decode(&s.State); err != nil {
		return nil, fmt.Errorf("invalid tf state file")
	}
	if s.State.Version != StateVersion {
		return nil, fmt.Errorf("unsupported state version %d", s.State.Version)
	}
	return &s, nil
}
