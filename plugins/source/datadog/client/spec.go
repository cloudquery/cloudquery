package client

import (
	_ "embed"
	"errors"
)

// Spec the (nested) spec used by the Datadog source plugin.
type Spec struct {
	// Specify which accounts to sync data from.
	Accounts []Account `json:"accounts" jsonschema:"required,minItems=1"`
	// A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency"`
	// The Datadog site to connect to. This is usually one of datadoghq.com or datadoghq.eu - see site documentation for more information.
	Site string `json:"site"`
}

// Account used to specify one or more accounts to extract information from.
type Account struct {
	// Account name.
	Name string `json:"name" jsonschema:"required,minLength=1"`
	// Datadog API key.
	APIKey string `json:"api_key" jsonschema:"required,minLength=1"`
	// Datadog App key.
	AppKey string `json:"app_key" jsonschema:"required,minLength=1"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 10000
	}
}

func (s *Spec) Validate() error {
	if len(s.Accounts) == 0 {
		return errors.New("no datadog accounts configured")
	}
	return nil
}

//go:embed schema.json
var JSONSchema string
