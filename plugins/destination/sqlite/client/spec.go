package client

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
}

func (*Spec) SetDefaults() {
	// stub for any future defaults
}
