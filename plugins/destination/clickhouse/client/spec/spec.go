package spec

import (
	"crypto/x509"
	_ "embed"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
)

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`
	Cluster          string `json:"cluster,omitempty"`

	Engine *Engine `json:"engine,omitempty"`

	CACert string `json:"ca_cert,omitempty"`

	BatchSize      int                  `json:"batch_size,omitempty" jsonschema:"minimum=1,default=10000"`
	BatchSizeBytes int                  `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=5242880"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) Options() (*clickhouse.Options, error) {
	options, err := clickhouse.ParseDSN(s.ConnectionString)
	if err != nil {
		return nil, err
	}

	// set database name to "default", if empty
	if len(options.Auth.Database) == 0 {
		options.Auth.Database = "default"
	}

	if tlsConfig := options.TLS; tlsConfig != nil && len(s.CACert) > 0 {
		if tlsConfig.RootCAs == nil {
			tlsConfig.RootCAs, err = x509.SystemCertPool()
			if err != nil {
				return nil, err
			}
		}

		if ok := tlsConfig.RootCAs.AppendCertsFromPEM([]byte(s.CACert)); !ok {
			return nil, fmt.Errorf("failed to append \"ca_cert\" value")
		}
	}

	return options, nil
}

func (s *Spec) SetDefaults() {
	if s.Engine == nil {
		s.Engine = DefaultEngine()
	}

	if s.BatchSize <= 0 {
		s.BatchSize = 10_000 // 10K
	}

	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = 5 << 20 // 5 MiB
	}

	if s.BatchTimeout == nil {
		d := configtype.NewDuration(20 * time.Second) // 20s
		s.BatchTimeout = &d
	}
}

func (s *Spec) Validate() error {
	return s.Engine.Validate()
}

// we need to set default for batch_timeout
func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "20s"
}

//go:embed schema.json
var JSONSchema string
