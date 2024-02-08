package spec

import _ "embed"

const (
	defaultConcurrency = 1000
)

type Spec struct {
	// Max number of requests per second to perform against the Hubspot API.
	MaxRequestsPerSecond *int `yaml:"max_requests_per_second,omitempty" json:"max_requests_per_second,omitempty" jsonschema:"minimum=1,default=5"`
	// Key-value map of options for each table. The key is the name of the table. The value is an options object.
	TableOptions TableOptions `yaml:"table_options,omitempty" json:"table_options,omitempty"`
	// Concurrency setting for the CloudQuery scheduler.
	Concurrency int `yaml:"concurrency,omitempty" json:"concurrency,omitempty" jsonschema:"minimum=1,default=1000"`
}

func (spec *Spec) SetDefaults() {
	// https://developers.hubspot.com/docs/api/usage-details#rate-limits
	// Hubspot, for Pro and Enterprise, accounts, has rate limits of:
	// - 15 requests / second / private-app
	// - 500,000 requests / day / org (5.7 requests / second / org).
	// I chose the default of 5, which should be safe for most accounts and use-cases (but may be too much for "Starter"
	// subscriptions in case cloudquery is run 24/7).
	var defaultRateLimitPerSecond = 5

	if spec.MaxRequestsPerSecond == nil || *spec.MaxRequestsPerSecond <= 0 {
		spec.MaxRequestsPerSecond = &defaultRateLimitPerSecond
	}

	if spec.Concurrency == 0 {
		spec.Concurrency = defaultConcurrency
	}
}

//go:embed schema.json
var JSONSchema string
