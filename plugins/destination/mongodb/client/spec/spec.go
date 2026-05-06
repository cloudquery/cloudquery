package spec

import (
	_ "embed"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4

	// Retries are opt-in. Default is a single attempt so behavior matches the
	// pre-write_retry plugin and we don't risk duplicate documents on
	// write_mode: append tables (see docs/overview.md callout).
	defaultWriteRetryMaxAttempts = 1
	defaultWriteRetryMaxBackoff  = 10 * time.Second
)

type Spec struct {
	// MongoDB URI as described in the official MongoDB [documentation](https://www.mongodb.com/docs/manual/reference/connection-string/).
	//
	// Example connection strings:
	// - `"mongodb://username:password@hostname:port/database"` basic connection
	// - `"mongodb+srv://username:password@cluster.example.com/database"` connecting to a MongoDB Atlas cluster
	// - `"mongodb://localhost:27017/myDatabase?authSource=admin"` specify authentication source
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1"`

	// Database to sync the data to.
	Database string `json:"database" jsonschema:"required,minLength=1"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`

	// Use AWS IAM credentials. If used this will override any credentials set in the connection_string
	AWSCredentials *Credentials `json:"aws_credentials,omitempty"`

	// Configures exponential-backoff retries around each write batch to absorb
	// transient MongoDB network errors (e.g. `write tcp ...: broken pipe`) that
	// are not covered by the driver's single built-in retry. Retries are
	// disabled by default (single attempt). Set `write_retry.max_attempts` >= 2
	// to enable. Read the duplicate-write caveat in the destination docs
	// before enabling for `write_mode: append` tables.
	WriteRetry *WriteRetryConfig `json:"write_retry,omitempty"`
}

type WriteRetryConfig struct {
	// Maximum number of write attempts per batch, including the initial attempt. Default is 1 (no retries).
	MaxAttempts int `json:"max_attempts,omitempty" jsonschema:"minimum=1,default=1"`

	// Maximum backoff between retry attempts.
	MaxBackoff *configtype.Duration `json:"max_backoff,omitempty" jsonschema:"default=10s"`

	// When `true`, each retried write batch runs inside a MongoDB
	// [transaction](https://www.mongodb.com/docs/manual/core/transactions/) so a
	// retry that follows a partially-applied write rolls back the partial state
	// before re-attempting. This eliminates the duplicate-document risk on
	// `write_mode: append` tables (where `_id` is server-generated and the
	// driver has no txnNumber to dedupe on).
	//
	// Requires a replica set, sharded cluster, or load-balanced deployment —
	// transactions are not supported on standalone MongoDB, and enabling this
	// against a standalone server will surface a driver error at write time.
	//
	// Has no effect when `max_attempts` is 1 (retries disabled).
	UseTransactions bool `json:"use_transactions,omitempty"`
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
	if s.WriteRetry == nil {
		s.WriteRetry = &WriteRetryConfig{}
	}
	s.WriteRetry.SetDefaults()
}

func (r *WriteRetryConfig) SetDefaults() {
	if r.MaxAttempts == 0 {
		r.MaxAttempts = defaultWriteRetryMaxAttempts
	}
	if r.MaxBackoff == nil {
		d := configtype.NewDuration(defaultWriteRetryMaxBackoff)
		r.MaxBackoff = &d
	}
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return errors.New("connection_string is required")
	}
	if s.Database == "" {
		return errors.New("database is required")
	}
	if s.AWSCredentials != nil {
		if (s.AWSCredentials.RoleARN != "" || s.AWSCredentials.RoleSessionName != "" || s.AWSCredentials.ExternalID != "" || s.AWSCredentials.LocalProfile != "") && s.AWSCredentials.Default {
			return errors.New("`default` cannot be used with any other credential options")
		}
		if s.AWSCredentials.RoleARN == "" && s.AWSCredentials.LocalProfile == "" && !s.AWSCredentials.Default {
			return errors.New("one of `role_arn`, `local_profile`, or `default` must be set")
		}
	}

	if s.WriteRetry != nil {
		if s.WriteRetry.MaxAttempts < 1 {
			return errors.New("`write_retry.max_attempts` must be >= 1")
		}
		if s.WriteRetry.MaxBackoff != nil && s.WriteRetry.MaxBackoff.Duration() < 0 {
			return errors.New("`write_retry.max_backoff` must be >= 0")
		}
	}

	return nil
}
