package client

import (
	_ "embed"
	"errors"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4 // 4MB
)

type Spec struct {
	// Name of the database and extra connection options, such as `my_db`.
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`

	// MotherDuck API token. If empty, the plugin will open a web browser to authenticate.
	Token string `json:"token,omitempty"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`

	// Enables debug logging
	Debug bool `json:"debug,omitempty" jsonschema:"default=false"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
}

func (s Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("connection_string is required")
	}
	if strings.HasPrefix(s.ConnectionString, "md:") {
		return errors.New("connection_string should not start with 'md:'")
	}

	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Not = &jsonschema.Schema{
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			properties := jsonschema.NewProperties()
			connStr := *sc.Properties.Value("connection_string")
			connStr.Pattern = "^md:.+"
			properties.Set("connection_string", &connStr)
			return properties
		}(),
	}
}
