package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type Spec struct {
	StreamARN  string `json:"stream_arn"`
	NoRotate   bool   `json:"no_rotate,omitempty"`
	MaxRetries *int   `json:"max_retries,omitempty"`

	MaxRecordSizeBytes int `json:"max_record_size_bytes,omitempty"`
	MaxBatchRecords    int `json:"max_batch_records,omitempty"`
	MaxBatchSizeBytes  int `json:"max_batch_size_bytes,omitempty"`
}

const (
	defaultMaxRecordSizeBytes = 1024000
	defaultMaxBatchRecords    = 500
	defaultMaxBatchSizeBytes  = 4194000
)

func (s *Spec) SetDefaults() {
	if s.MaxRetries == nil {
		s.MaxRetries = new(int)
		*s.MaxRetries = 5
	}
	if s.MaxRecordSizeBytes < 1 {
		s.MaxRecordSizeBytes = defaultMaxRecordSizeBytes
	}
	if s.MaxBatchRecords < 1 {
		s.MaxBatchRecords = defaultMaxBatchRecords
	}
	if s.MaxBatchSizeBytes < 1 {
		s.MaxBatchSizeBytes = defaultMaxBatchSizeBytes
	}
}

func (s *Spec) Validate() error {
	if s.StreamARN == "" {
		return fmt.Errorf("kinesis firehose Stream ARN is required")
	}
	parsedARN, err := arn.Parse(s.StreamARN)
	if err != nil {
		return fmt.Errorf("kinesis firehose Stream ARN is invalid")
	}
	if parsedARN.Service != "firehose" {
		return fmt.Errorf("kinesis firehose Stream ARN is invalid")
	}
	return nil
}
