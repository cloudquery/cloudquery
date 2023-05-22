package client

type Spec struct {
	Subscriptions        []string `json:"subscriptions"`
	CloudName            string   `json:"cloud_name"`
	DiscoveryConcurrency int      `json:"discovery_concurrency"`
	SkipSubscriptions    []string `json:"skip_subscriptions"`
}

func (s *Spec) SetDefaults() {
	if s.DiscoveryConcurrency <= 0 {
		s.DiscoveryConcurrency = 400
	}
}
