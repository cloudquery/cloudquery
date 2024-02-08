package client

import _ "embed"

type Spec struct {
	// Path to a file, such as `./mydb.sql`
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`
}

//go:embed schema.json
var JSONSchema string

func (*Spec) SetDefaults() {
	// stub for any future defaults
}
