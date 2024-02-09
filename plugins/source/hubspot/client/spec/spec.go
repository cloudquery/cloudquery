package spec

import (
	_ "embed"
	"fmt"
	"os"
)

const (
	defaultConcurrency = 1000
)

type Spec struct {
	// In order for CloudQuery to sync resources from your HubSpot setup, you will need to authenticate with your HubSpot account. You will need to create a [HubSpot Private App](https://developers.hubspot.com/docs/api/private-apps), and copy the App Token here.
	// If not specified `HUBSPOT_APP_TOKEN` environment variable will be used instead.
	AppToken string `json:"app_token,omitempty" jsonschema:"minLength=1"`
	// Max number of requests per second to perform against the Hubspot API.
	MaxRequestsPerSecond int `json:"max_requests_per_second,omitempty" jsonschema:"minimum=1,default=5"`
	// Key-value map of options for each table. The key is the name of the table. The value is an options object.
	TableOptions TableOptions `json:"table_options,omitempty"`
	// Concurrency setting for the CloudQuery scheduler.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=1000"`
}

func (s *Spec) SetDefaults() {
	if len(s.AppToken) == 0 {
		s.AppToken = os.Getenv("HUBSPOT_APP_TOKEN")
	}

	// https://developers.hubspot.com/docs/api/usage-details#rate-limits
	// Hubspot, for Pro and Enterprise, accounts, has rate limits of:
	// - 15 requests / second / private-app
	// - 500,000 requests / day / org (5.7 requests / second / org).
	// I chose the default of 5, which should be safe for most accounts and use-cases (but may be too much for "Starter"
	// subscriptions in case cloudquery is run 24/7).
	var defaultRateLimitPerSecond = 5

	if s.MaxRequestsPerSecond <= 0 {
		s.MaxRequestsPerSecond = defaultRateLimitPerSecond
	}

	if s.Concurrency <= 0 {
		s.Concurrency = defaultConcurrency
	}
}

func (s Spec) Validate() error {
	if s.AppToken == "" {
		return fmt.Errorf("app_token is required")
	}
	return nil
}

//go:embed schema.json
var JSONSchema string
