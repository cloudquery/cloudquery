package spec

type Spec struct {
	Subscriptions        []string `json:"subscriptions" jsonschema:"minLength=1,uniqueItems=true"`
	CloudName            string   `json:"cloud_name"`
	SkipSubscriptions    []string `json:"skip_subscriptions" jsonschema:"minLength=1,uniqueItems=true"`
	NormalizeIDs         bool     `json:"normalize_ids"`
	OIDCToken            string   `json:"oidc_token"`
	DiscoveryConcurrency int      `json:"discovery_concurrency" jsonschema:"minimum=1,default=400"`
	Concurrency          int      `json:"concurrency" jsonschema:"minimum=1,default=50000"`
}

func (s *Spec) SetDefaults() {
	if s.DiscoveryConcurrency <= 0 {
		s.DiscoveryConcurrency = 400
	}
	if s.Concurrency <= 0 {
		const defaultConcurrency = 50000
		s.Concurrency = defaultConcurrency
	}
}
