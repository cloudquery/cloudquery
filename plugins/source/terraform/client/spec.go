package client

import (
	"errors"
)

type Spec struct {
	Backends    []BackendConfigBlock `json:"backends,omitempty"`
	Concurrency int                  `json:"concurrency,omitempty"`
}

func (s *Spec) Validate() error {
	if len(s.Backends) == 0 {
		return errors.New("no backends were provided")
	}
	for _, bc := range s.Backends {
		if err := bc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
