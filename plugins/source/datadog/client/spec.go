package client

import _ "embed"

// Spec the (nested) spec used by the Datadog source plugin.
type Spec struct {
	// Specify which accounts to sync data from.
	Accounts []Account `json:"accounts"`
	// A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency"`
	// The Datadog site to connect to. This is usually one of datadoghq.com or datadoghq.eu - see site documentation for more information.
	Site string `json:"site"`
}

// Account used to specify one or more accounts to extract information from.
type Account struct {
	// Account name.
	Name string `json:"name"`
	// Datadog API key.
	APIKey string `json:"api_key"`
	// Datadog App key.
	AppKey string `json:"app_key"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 10000
	}
}

//go:embed schema.json
var JSONSchema string
