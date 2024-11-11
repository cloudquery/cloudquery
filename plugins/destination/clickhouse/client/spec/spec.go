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

// CloudQuery ClickHouse destination plugin spec.
type Spec struct {
	// Connection string to connect to the database.
	// See [SDK documentation](https://github.com/ClickHouse/clickhouse-go#dsn) for more details.
	//
	// Example connection string:
	//
	// - `"clickhouse://username:password@host1:9000,host2:9000/database?dial_timeout=200ms&max_execution_time=60"`
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`

	// Cluster name to be used for [distributed DDL](https://clickhouse.com/docs/en/sql-reference/distributed-ddl).
	// If the value is empty, DDL operations will affect only the server the plugin is connected to.
	Cluster string `json:"cluster,omitempty"`

	// Engine to be used for tables.
	// Only [`*MergeTree` family](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family) is supported at the moment.
	Engine *Engine `json:"engine,omitempty"`

	// PEM-encoded certificate authorities.
	// When set, a certificate pool will be created by appending the certificates to the system pool.
	//
	// See [file variable substitution](/docs/advanced-topics/environment-variable-substitution#file-variable-substitution-example)
	// for how to read this value from a file.
	CACert string `json:"ca_cert,omitempty"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=5242880"`

	// Maximum interval between batch writes.
	BatchTimeout *configtype.Duration `json:"batch_timeout,omitempty"`

	// If true, the plugin will partition the data by the _cq_sync_group_id column.
	PartitionBySyncGroupID bool `json:"partition_by_sync_group_id,omitempty"`
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
