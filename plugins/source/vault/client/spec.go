package client

import "errors"

type Spec struct {
	VaultAddress string `json:"vault_address,omitempty"`
}

func (*Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	if s.VaultAddress == "" {
		return errors.New("no vault address provided")
	}
	return nil
}
