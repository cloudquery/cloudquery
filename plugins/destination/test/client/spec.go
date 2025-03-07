package client

import (
	_ "embed"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	// If true, will return an error on any write message rather than consume from the channel
	ErrorOnWrite bool `json:"error_on_write,omitempty" jsonschema:"default=false"`

	// If true, will return an error on migrate table messages rather than consume from the channel
	ErrorOnMigrate bool `json:"error_on_migrate,omitempty" jsonschema:"default=false"`

	// If true, will return an error on insert record messages rather than consume from the channel
	ErrorOnInsert bool `json:"error_on_insert,omitempty" jsonschema:"default=false"`

	// If true, will call os.Exit(1) on any write message rather than consume from the channel
	ExitOnWrite bool `json:"exit_on_write,omitempty" jsonschema:"default=false"`

	// If true, will call os.Exit(1) on migrate table messages rather than consume from the channel
	ExitOnMigrate bool `json:"exit_on_migrate,omitempty" jsonschema:"default=false"`

	// If true, will call os.Exit(1) on insert record messages rather than consume from the channel
	ExitOnInsert bool `json:"exit_on_insert,omitempty" jsonschema:"default=false"`

	// Whether to use a BatchWriter or not. Only one BatchWriter option should be true.
	BatchWriter bool `json:"batch_writer" jsonschema:"default=false"`

	// Whether to use a StreamBatchWriter or not. Only one BatchWriter option should be true.
	StreamBatchWriter bool `json:"stream_batch_writer" jsonschema:"default=false"`

	// Whether to use a MixedBatchWriter or not. Only one BatchWriter option should be true.
	MixedBatchWriter bool `json:"mixed_batch_writer" jsonschema:"default=false"`

	// Maximum number of items that may be grouped together to be written in a single write.
	//
	// Defaults to `10000`.
	BatchSize *int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	//
	// Defaults to `52428800` (50 MiB).
	BatchSizeBytes *int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=52428800"`

	// Maximum interval between batch writes.
	//
	// Defaults to `30s`.
	BatchTimeout *configtype.Duration `json:"batch_timeout" jsonschema:"default=30s"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.BatchSize == nil {
		s.BatchSize = ptr(int64(10000))
	}
	if s.BatchSizeBytes == nil {
		s.BatchSizeBytes = ptr(int64(50 * 1024 * 1024)) // 50 MiB
	}
	if s.BatchTimeout == nil {
		s.BatchTimeout = ptr(configtype.NewDuration(30 * time.Second))
	}
}

func (s *Spec) Validate() error {
	if s.BatchWriter && s.StreamBatchWriter {
		return errors.New("batch_writer and stream_batch_writer cannot be true at the same time")
	}
	if s.BatchWriter && s.MixedBatchWriter {
		return errors.New("batch_writer and mixed_batch_writer cannot be true at the same time")
	}
	if s.StreamBatchWriter && s.MixedBatchWriter {
		return errors.New("stream_batch_writer and mixed_batch_writer cannot be true at the same time")
	}
	return nil
}

func ptr[A any](a A) *A {
	return &a
}
