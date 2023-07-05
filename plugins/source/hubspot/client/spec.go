package client

const (
	defaultConcurrency = 1000
)

type Spec struct {
	MaxRequestsPerSecond *int         `yaml:"max_requests_per_second,omitempty" json:"max_requests_per_second,omitempty"`
	TableOptions         TableOptions `yaml:"table_options,omitempty" json:"table_options,omitempty"`
	Concurrency          int          `yaml:"concurrency,omitempty" json:"concurrency,omitempty"`
}

type TableOptions map[string]*TableOptionsSpec

type TableOptionsSpec struct {
	Properties   []string `yaml:"properties,omitempty" json:"properties,omitempty"`
	Associations []string `yaml:"associations,omitempty" json:"associations,omitempty"`
}

func (spec *Spec) SetDefaults() {
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

	if spec.Concurrency == 0 {
		spec.Concurrency = defaultConcurrency
	}
}

func (ts TableOptions) ForTable(name string) *TableOptionsSpec {
	return ts[name]
}

func (to *TableOptionsSpec) GetProperties() []string {
	if to == nil {
		return nil
	}
	return to.Properties
}

func (to *TableOptionsSpec) GetAssociations() []string {
	if to == nil {
		return nil
	}
	return to.Associations
}
