package client

type Spec struct {
	Endpoint string `json:"endpoint,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Endpoint == "" {
		s.Endpoint = "https://pricing.us-east-1.amazonaws.com"
	}
}
