package client

const defaultConcurrency = 1000

type Spec struct {
	// Used in API requests to filter only resources related to these team ids.
	// Used in the tables: ["escalation_policies", "incidents", "maintenance_windows", "services", "users"]
	TeamIds []string `json:"team_ids"`

	MaxRequestsPerSecond *int `json:"max_requests_per_second"`

	Concurrency int `json:"concurrency"`
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
