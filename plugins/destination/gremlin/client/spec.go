package client

import (
	"fmt"
	"runtime"
	"strings"
)

type Spec struct {
	Endpoint string `json:"endpoint"`
	Insecure bool   `json:"insecure"`

	AuthMode authMode `json:"auth_mode"`

	// Static credentials
	Username string `json:"username"`
	Password string `json:"password"`

	// Backoff
	MaxRetries int `json:"max_retries"`

	// AWS specific settings
	AWSRegion string `json:"aws_region"`

	// Connection settings
	MaxConcurrentConnections int `json:"max_concurrent_connections"`

	// Whether to use all Gremlin types or just a basic subset
	CompleteTypes bool `json:"complete_types"`

	BatchSize      int `json:"batch_size"`
	BatchSizeBytes int `json:"batch_size_bytes"`
}

type authMode string

const (
	authModeNone  = authMode("none")
	authModeBasic = authMode("basic")
	authModeAWS   = authMode("aws")
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
		s.AuthMode = authModeNone
	} else {
		s.AuthMode = authMode(strings.ToLower(string(s.AuthMode)))
	}

	if s.MaxRetries <= 0 {
		s.MaxRetries = 5 // 5 retries by default
	}

	if s.MaxConcurrentConnections <= 0 {
		s.MaxConcurrentConnections = runtime.NumCPU()
	}

	if s.BatchSize <= 0 {
		s.BatchSize = 200
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = 1024 * 1024 * 4
	}
}

func (s *Spec) Validate() error {
	if s.Endpoint == "" {
		return fmt.Errorf("endpoint is required")
	}
	if s.AuthMode != authModeNone && s.AuthMode != authModeBasic && s.AuthMode != authModeAWS {
		return fmt.Errorf("invalid auth_mode, valid values are %q, %q and %q", authModeNone, authModeBasic, authModeAWS)
	}
	if s.AuthMode == authModeAWS && s.AWSRegion == "" {
		return fmt.Errorf("aws_region is required when auth_mode is %q", authModeAWS)
	}
	if s.AuthMode == authModeNone && (s.Username != "" || s.Password != "") {
		return fmt.Errorf("username or password specified with auth_mode %q. Set auth mode to %q or remove username and password", authModeNone, authModeBasic)
	}

	return nil
}
