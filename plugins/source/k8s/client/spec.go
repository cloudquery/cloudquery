package client

const (
	defaultConcurrency = 50000
)

type Spec struct {
	Contexts    []string `yaml:"contexts,omitempty" json:"contexts"`
	Concurrency int      `yaml:"concurrency,omitempty" json:"concurrency"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}
