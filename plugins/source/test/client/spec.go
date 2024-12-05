package client

import (
	_ "embed"
	"fmt"
	"strings"
)

type Spec struct {
	// Number of clients to create
	NumClients int `json:"num_clients" jsonschema:"minimum=1,default=1"`

	// Number of rows to generate in test_some_table.
	NumRows *int `json:"num_rows" jsonschema:"minimum=0,default=1"`

	// Number of rows to generate (per row of parent) in test_sub_table.
	NumSubRows *int `json:"num_sub_rows" jsonschema:"minimum=0,default=10"`

	// Number of `extra_column_*` columns to generate per row in test_sub_table. The type of the columns is int64.
	NumSubCols *int `json:"num_sub_cols" jsonschema:"minimum=0,default=1"`

	// Required environment variables. The plugin will fail if these are not set
	// to the correct values. Specified in `key=value` format. Use `key=` to specify
	// that the environment variable should be not set or empty.
	RequiredEnv []string `json:"required_env" jsonschema:"pattern=^[\\w]+\\=[\\w]*$"`

	// If true, the plugin will fail immediately at the table resolver level, before any resources are synced
	FailImmediately bool `json:"fail_immediately" jsonschema:"default=false"`
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
	if s.NumSubCols == nil || *s.NumSubCols < 0 {
		i := 1
		s.NumSubCols = &i
	}
	if s.RequiredEnv == nil {
		s.RequiredEnv = []string{}
	}
}

func (s *Spec) Validate() error {
	for _, v := range s.RequiredEnv {
		if strings.Count(v, "=") != 1 {
			return fmt.Errorf("required_env must be in the format `key=value`")
		}
	}
	return nil
}
