package client

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
}

func (s *Spec) SetDefaults() {}
