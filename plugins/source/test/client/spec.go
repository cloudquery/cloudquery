package client

type Spec struct {
	NumClients int `json:"num_clients"`
}

func (s *Spec) SetDefaults() {
	if s.NumClients <= 0 {
		s.NumClients = 1
	}
}

func (*Spec) Validate() error {
	return nil
}
