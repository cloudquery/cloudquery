package client

import (
	"errors"
	"fmt"
	"net/url"
)

const defaultConcurrency = 10000

type Spec struct {
	VaultAddress string `json:"vault_address,omitempty"`
	Concurrency  int    `json:"concurrency"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency <= 0 {
		s.Concurrency = defaultConcurrency
	}
}

func (s *Spec) Validate() error {
	if s.VaultAddress == "" {
		return errors.New("no vault address provided")
	}
	if _, err := url.ParseRequestURI(s.VaultAddress); err != nil {
		return fmt.Errorf("invalid vault address provided %q: %w", s.VaultAddress, err)
	}
	return nil
}
