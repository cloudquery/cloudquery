package client

import _ "embed"

type Spec struct {
	// Number of clients to create
	NumClients int `json:"num_clients" jsonschema:"minimum=1,default=1"`

	// Number of rows to generate in test_some_table.
	NumRows *int `json:"num_rows" jsonschema:"minimum=0,default=1"`

	// Number of rows to generate (per row of parent) in test_sub_table.
	NumSubRows *int `json:"num_sub_rows" jsonschema:"minimum=0,default=10"`
}

//go:embed schema.json
var JSONSchema string

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
