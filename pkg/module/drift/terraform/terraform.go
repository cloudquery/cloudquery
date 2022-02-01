package terraform

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// LoadState loads the given reader into tfstate
func LoadState(reader io.Reader) (*Data, error) {
	var s Data
	if err := json.NewDecoder(reader).Decode(&s.State); err != nil {
		return nil, fmt.Errorf("invalid tf state file: %w", err)
	}
	return &s, nil
}

// ValidateStateVersion validates the given tfstate version to be version 4
func ValidateStateVersion(s *Data) (bool, error) {
	if s.State.Version == nil {
		return true, errors.New("unspecified tfstate version, allowing")
	}
	switch v := string(*s.State.Version); v {
	case "4":
		return true, nil
	case "2", "3":
		return false, fmt.Errorf("unsupported tfstate version %s", v)
	default:
		return true, fmt.Errorf("unknown tfstate version %s, allowing", v)
	}
}
