package client

const (
	defaultConcurrency = 50000
)

type Spec struct {
	Subscriptions        []string `json:"subscriptions"`
	CloudName            string   `json:"cloud_name"`
	DiscoveryConcurrency int      `json:"discovery_concurrency"`
	SkipSubscriptions    []string `json:"skip_subscriptions"`
	NormalizeIDs         bool     `json:"normalize_ids"`
	Concurrency          int      `json:"concurrency"`
}

func (s *Spec) SetDefaults() {
	if s.DiscoveryConcurrency <= 0 {
		s.DiscoveryConcurrency = 400
	}
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}
