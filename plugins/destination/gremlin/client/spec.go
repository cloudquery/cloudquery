package client

import (
	"fmt"
	"strings"
)

type Spec struct {
	Endpoint string `json:"endpoint"`
	Insecure bool   `json:"insecure"`

	AuthMode authMode `json:"auth_mode"`

	// Static credentials
	Username string `json:"username"`
	Password string `json:"password"`
}

type authMode string

const (
	authModeBasic = authMode("basic")
	authModeIAM   = authMode("iam")
)

func (s *Spec) SetDefaults() {
	if s.Endpoint != "" {
		// Default to "wss://<endpoint>:8182" where "wss://" and ":8182" are optional
		e := strings.SplitN(s.Endpoint, "://", 2)
		if len(e) == 1 {
			e = []string{"wss", e[0]}
		}
		if !strings.Contains(e[1], ":") {
			e[1] += ":8182"
		}
		s.Endpoint = strings.Join(e, "://")
	}

	if s.AuthMode == "" {
		s.AuthMode = authModeBasic
	}
}

func (s *Spec) Validate() error {
	if s.Endpoint == "" {
		return fmt.Errorf("endpoint is required")
	}
	if s.AuthMode != authModeBasic && s.AuthMode != authModeIAM {
		return fmt.Errorf("invalid auth_mode, valid values are %q and %q", authModeBasic, authModeIAM)
	}
	return nil
}
