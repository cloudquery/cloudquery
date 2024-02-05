package client

import _ "embed"

const defaultConcurrency = 1000

type Spec struct {
	// Used in API requests to filter only resources related to these team ids.
	// Used in the tables: ["escalation_policies", "incidents", "maintenance_windows", "services", "users"]
	TeamIds []string `json:"team_ids"`

	// PagerDuty API is heavily rate-limited (900 requests/min = 15 requests/sec, across the entire organization).
	// This option allows you to control the rate at which the plugin will make requests to the API.
	// You can reduce this parameter in case you are still seeing rate limit errors (status code 429), or increase
	// it if your PagerDuty API quota is higher. See https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUz-rate-limiting#what-are-our-limits for more info.
	MaxRequestsPerSecond *int `json:"max_requests_per_second" jsonschema:"minimum=1,default=10"`

	// A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=1000"`
}

func (spec *Spec) SetDefaults() {
	// Calculated as 66% of 900 requests per minute.
	// https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUz-rate-limiting#what-are-our-limits
	// The 900 requests per minute is for the entire oragnization, so we don't actually want to come too
	// close to it.
	var defaultRateLimitPerSecond = 10

	if spec.MaxRequestsPerSecond == nil || *spec.MaxRequestsPerSecond == 0 {
		spec.MaxRequestsPerSecond = &defaultRateLimitPerSecond
	}

	if spec.Concurrency <= 0 {
		spec.Concurrency = defaultConcurrency
	}
}

func (*Spec) Validate() error {
	return nil
}

//go:embed schema.json
var JSONSchema string
