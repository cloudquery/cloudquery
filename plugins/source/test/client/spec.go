package client

type Spec struct {
	NumClients int  `json:"num_clients"`  // Number of clients to create. Default: 1
	NumRows    *int `json:"num_rows"`     // Number of rows to generate in test_some_table. Defaul: 1
	NumSubRows *int `json:"num_sub_rows"` // Number of rows to generate (per row of parent) in test_sub_table. Default: 10
}

func (s *Spec) SetDefaults() {
	if s.NumClients <= 0 {
		s.NumClients = 1
	}
	if s.NumRows == nil || *s.NumRows < 0 {
		i := 1
		s.NumRows = &i
	}
	if s.NumSubRows == nil || *s.NumSubRows < 0 {
		i := 10
		s.NumSubRows = &i
	}
}

func (*Spec) Validate() error {
	return nil
}
