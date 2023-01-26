package client

type Spec struct {
	// Used in API requests to filter only resources related to these team ids.
	// Used in the tables: ["escalation_policies", "incidents", "maintenance_windows", "services", "users"]
	TeamIds []string `yaml:"team_ids,omitempty" json:"team_ids"`

	MaxRequestsPerSecond *int `yaml:"max_requests_per_second,omitempty" json:"max_requests_per_second,omitempty"`
}

func (spec *Spec) setDefaults() {
	// Calculated as 66% of 900 requests per minute.
	// https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTUz-rate-limiting#what-are-our-limits
	// The 900 requests per minute is for the entire oragnization, so we don't actually want to come too
	// close to it.
	var defaultRateLimitPerSecond = 10

	if spec.MaxRequestsPerSecond == nil || *spec.MaxRequestsPerSecond == 0 {
		spec.MaxRequestsPerSecond = &defaultRateLimitPerSecond
	}
}
