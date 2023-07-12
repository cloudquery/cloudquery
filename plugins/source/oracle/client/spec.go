package client

type Spec struct {
	Concurrency int `json:"concurrency,omitempty"`
}

func (*Spec) Validate() error {
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
