package spec

import (
	_ "embed"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

// Amazon Kinesis Firehose destination plugin spec.
type Spec struct {
	// Kinesis Firehose delivery stream ARN where data will be sent to.
	// Format: `arn:${Partition}:firehose:${Region}:${Account}:deliverystream/${DeliveryStreamName}`.
	StreamARN string `json:"stream_arn" jsonschema:"required,minLength=1"`

	// Amount of retries to perform when writing a batch.
	MaxRetries int `json:"max_retries,omitempty" jsonschema:"minimum=1,default=5"`

	// Number of bytes (as Arrow buffer size) to write before starting a new record.
	MaxRecordSizeBytes int `json:"max_record_size_bytes,omitempty" jsonschema:"minimum=1,default=1024000"`

	// Number of records allowed in a single batch.
	MaxBatchRecords int `json:"max_batch_records,omitempty" jsonschema:"minimum=1,default=500"`

	// Number of bytes allowed in a single batch.
	MaxBatchSizeBytes int `json:"max_batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194000"`
}

func (s *Spec) SetDefaults() {
	if s.MaxRetries < 1 {
		s.MaxRetries = 5
	}
	if s.MaxRecordSizeBytes < 1 {
		const defaultMaxRecordSizeBytes = 1024000
		s.MaxRecordSizeBytes = defaultMaxRecordSizeBytes
	}
	if s.MaxBatchRecords < 1 {
		const defaultMaxBatchRecords = 500
		s.MaxBatchRecords = defaultMaxBatchRecords
	}
	if s.MaxBatchSizeBytes < 1 {
		const defaultMaxBatchSizeBytes = 4194000
		s.MaxBatchSizeBytes = defaultMaxBatchSizeBytes
	}
}

func (s *Spec) Validate() error {
	if len(s.StreamARN) == 0 {
		return errors.New("kinesis firehose Stream ARN is required")
	}
	parsedARN, err := arn.Parse(s.StreamARN)
	if err != nil {
		return errors.New("kinesis firehose Stream ARN is invalid")
	}
	if parsedARN.Service != "firehose" {
		return errors.New("kinesis firehose Stream ARN is invalid")
	}
	return nil
}

//go:embed schema.json
var JSONSchema string
