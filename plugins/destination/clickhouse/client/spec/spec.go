package spec

import (
	"crypto/x509"
	_ "embed"
	"errors"
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

	// Enables partitioning of tables via the `PARTITION BY` clause.
	Partition []PartitionStrategy `json:"partition,omitempty"`

	// Enables setting table sort keys via the `ORDER BY` clause.
	OrderBy []OrderByStrategy `json:"order,omitempty"`
}

type PartitionStrategy struct {
	// Table glob patterns that apply for this partitioning.
	//
	// If unset, the partitioning will apply to all tables.
	//
	// If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.
	//
	// Partition strategy table patterns should be disjointed sets: if a table matches two partition strategies,
	// an error will be raised at runtime.
	Tables []string `json:"tables,omitempty"`

	// Table glob patterns that should be skipped for this partitioning.
	//
	// If unset, no tables will be skipped.
	//
	// If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.
	//
	// Partition strategy table patterns should be disjointed sets: if a table matches two partition strategies,
	// an error will be raised at runtime.
	SkipTables []string `json:"skip_tables,omitempty"`

	// Partitioning strategy to use, e.g. `toYYYYMM(_cq_sync_time)`,
	// the string is passed as is after "PARTITION BY" clause with no validation or quoting.
	//
	// An unset partition_by is not valid.
	PartitionBy string `json:"partition_by"`

	// Skip incremental tables from partitioning.
	SkipIncrementalTables bool `json:"skip_incremental_tables,omitempty"`
}

type OrderByStrategy struct {
	// Table glob patterns that apply for this ORDER BY clause.
	//
	// If unset, the ORDER BY clause will apply to all tables.
	//
	// If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.
	//
	// Order by strategy table patterns should be disjointed sets: if a table matches two order by strategies,
	// an error will be raised at runtime.
	Tables []string `json:"tables,omitempty"`

	// Table glob patterns that should be skipped for this ORDER BY clause.
	//
	// If unset, no tables will be skipped.
	//
	// If a table matches both a pattern in `tables` and `skip_tables`, the table will be skipped.
	//
	// Order by strategy table patterns should be disjointed sets: if a table matches two order by strategies,
	// an error will be raised at runtime.
	SkipTables []string `json:"skip_tables,omitempty"`

	// ORDER BY list of expressions to use, e.g. `_cq_sync_group_id, toYYYYMM(_cq_sync_time), _cq_id`,
	// the strings are passed as is after "ORDER BY" clause, separated by commas, with no validation or quoting.
	//
	// An unset order_by is not valid.
	OrderBy []string `json:"order_by"`
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
			return nil, errors.New("failed to append \"ca_cert\" value")
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

	for i, p := range s.Partition {
		if len(p.Tables) == 0 {
			s.Partition[i].Tables = []string{"*"}
		}
	}

	for i, o := range s.OrderBy {
		if len(o.Tables) == 0 {
			s.OrderBy[i].Tables = []string{"*"}
		}
	}
}

func (s *Spec) Validate() error {
	for _, p := range s.Partition {
		if len(p.PartitionBy) == 0 {
			return errors.New("partition_by is required")
		}
	}

	for _, o := range s.OrderBy {
		if len(o.OrderBy) == 0 {
			return errors.New("order_by is required")
		}
	}

	return s.Engine.Validate()
}

// we need to set default for batch_timeout
func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "20s"
}

//go:embed schema.json
var JSONSchema string
