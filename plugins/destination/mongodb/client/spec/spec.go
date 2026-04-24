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

	defaultWriteRetryMaxAttempts    = 5
	defaultWriteRetryInitialBackoff = 500 * time.Millisecond
	defaultWriteRetryMaxBackoff     = 10 * time.Second
	defaultWriteRetryMaxElapsed     = 30 * time.Second
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
	// are not covered by the driver's single built-in retry. Omit to use the
	// defaults (5 attempts, 500ms initial backoff, 10s max backoff, 30s max
	// total elapsed time).
	WriteRetry *WriteRetryConfig `json:"write_retry,omitempty"`
}

type WriteRetryConfig struct {
	// Maximum number of write attempts per batch, including the initial attempt. Set to 1 to disable retries.
	MaxAttempts int `json:"max_attempts,omitempty" jsonschema:"minimum=1,default=5"`

	// Initial backoff between retry attempts. Grows exponentially up to `max_backoff`.
	InitialBackoff *configtype.Duration `json:"initial_backoff,omitempty" jsonschema:"default=500ms"`

	// Maximum backoff between retry attempts.
	MaxBackoff *configtype.Duration `json:"max_backoff,omitempty" jsonschema:"default=10s"`

	// Maximum total time to spend retrying a single write batch before giving up.
	MaxElapsed *configtype.Duration `json:"max_elapsed,omitempty" jsonschema:"default=30s"`
}

func (r *WriteRetryConfig) GetMaxAttempts() int {
	if r == nil || r.MaxAttempts <= 0 {
		return defaultWriteRetryMaxAttempts
	}
	return r.MaxAttempts
}

func (r *WriteRetryConfig) GetInitialBackoff() time.Duration {
	if r == nil || r.InitialBackoff == nil || r.InitialBackoff.Duration() <= 0 {
		return defaultWriteRetryInitialBackoff
	}
	return r.InitialBackoff.Duration()
}

func (r *WriteRetryConfig) GetMaxBackoff() time.Duration {
	if r == nil || r.MaxBackoff == nil || r.MaxBackoff.Duration() <= 0 {
		return defaultWriteRetryMaxBackoff
	}
	return r.MaxBackoff.Duration()
}

func (r *WriteRetryConfig) GetMaxElapsed() time.Duration {
	if r == nil || r.MaxElapsed == nil || r.MaxElapsed.Duration() <= 0 {
		return defaultWriteRetryMaxElapsed
	}
	return r.MaxElapsed.Duration()
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
		if s.WriteRetry.MaxAttempts < 0 {
			return errors.New("`write_retry.max_attempts` must be >= 1")
		}
		if s.WriteRetry.InitialBackoff != nil && s.WriteRetry.InitialBackoff.Duration() < 0 {
			return errors.New("`write_retry.initial_backoff` must be >= 0")
		}
		if s.WriteRetry.MaxBackoff != nil && s.WriteRetry.MaxBackoff.Duration() < 0 {
			return errors.New("`write_retry.max_backoff` must be >= 0")
		}
		if s.WriteRetry.MaxElapsed != nil && s.WriteRetry.MaxElapsed.Duration() < 0 {
			return errors.New("`write_retry.max_elapsed` must be >= 0")
		}
	}

	return nil
}
