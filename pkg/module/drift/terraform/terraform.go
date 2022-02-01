package terraform

import (
	"bytes"
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
	if bytes.Equal(*s.State.Version, []byte{'4'}) {
		return true, nil
	}
	if bytes.Equal(*s.State.Version, []byte{'2'}) || bytes.Equal(*s.State.Version, []byte{'3'}) {
		return false, fmt.Errorf("unsupported tfstate version %s", string(*s.State.Version))
	}
	return true, fmt.Errorf("unknown tfstate version %s, allowing", string(*s.State.Version))
}
