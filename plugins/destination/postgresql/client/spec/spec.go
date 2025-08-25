package spec

import (
	_ "embed"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	"github.com/jackc/pgx/v5/tracelog"
)

const (
	defaultBatchSize      = 10000
	defaultBatchSizeBytes = 100000000
	defaultBatchTimeout   = 60 * time.Second
)

type Spec struct {
	// Connection string to connect to the database. This can be a DSN or a URI, for example:
	//
	// - `"user=user password=pass host=localhost port=5432 dbname=mydb sslmode=disable"` DSN format
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=prefer"` connect with tcp and prefer TLS
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=disable&search_path=myschema"` connect with tcp, disable TLS and use a custom schema
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1,example=${POSTGRESQL_CONNECTION_STRING}"`

	// Available: `error`, `warn`, `info`, `debug`, `trace`.
	// Defines what [`pgx`](https://github.com/jackc/pgx) call events should be logged.
	PgxLogLevel LogLevel `json:"pgx_log_level,omitempty" jsonschema:"default=error"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=100000000"`

	// Maximum interval between batch writes.
	BatchTimeout configtype.Duration `json:"batch_timeout,omitempty"`

	// Option to create specific indexes to improve deletion performance
	CreatePerformanceIndexes bool `json:"create_performance_indexes,omitempty" jsonschema:"default=false"`

	RetryOnDeadlock bool `json:"retry_on_deadlock,omitempty" jsonschema:"default=false"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize <= 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.BatchTimeout.Duration() <= 0 {
		s.BatchTimeout = configtype.NewDuration(defaultBatchTimeout)
	}
	if s.PgxLogLevel == 0 {
		s.PgxLogLevel = LogLevel(tracelog.LogLevelError)
	}
}

func (s *Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("`connection_string` is required")
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Value("batch_timeout").Default = "60s"
}

//go:embed schema.json
var JSONSchema string
