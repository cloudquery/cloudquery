package client

type Spec struct {
	MaxRequestsPerSecond *int         `yaml:"max_requests_per_second,omitempty" json:"max_requests_per_second,omitempty"`
	Companies            TableOptions `yaml:"companies,omitempty" json:"companies,omitempty"`
	Contacts             TableOptions `yaml:"contacts,omitempty" json:"contacts,omitempty"`
	Deals                TableOptions `yaml:"deals,omitempty" json:"deals,omitempty"`
	LineItems            TableOptions `yaml:"line_items,omitempty" json:"line_items,omitempty"`
	Products             TableOptions `yaml:"products,omitempty" json:"products,omitempty"`
	Quotes               TableOptions `yaml:"quotes,omitempty" json:"quotes,omitempty"`
	Tickets              TableOptions `yaml:"tickets,omitempty" json:"tickets,omitempty"`
}

type TableOptions struct {
	Properties []string `yaml:"properties,omitempty" json:"properties,omitempty"`
}

func (spec *Spec) setDefaults() {
	// https://developers.hubspot.com/docs/api/usage-details#rate-limits
	// Hubspot, for Pro and Enterprise, accounts, has rate limits of:
	// - 15 requests / second / private-app
	// - 500,000 requests / day / org (5.7 requests / second / org).
	// I chose the default of 5, which should be safe for most accounts and use-cases (but may be too much for "Starter"
	// subscriptions in case cloudquery is run 24/7).
	var defaultRateLimitPerSecond = 5

	if spec.MaxRequestsPerSecond == nil || *spec.MaxRequestsPerSecond == 0 {
		spec.MaxRequestsPerSecond = &defaultRateLimitPerSecond
	}
}
