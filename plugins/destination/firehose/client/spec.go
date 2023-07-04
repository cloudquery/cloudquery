package client

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type Spec struct {
	StreamARN  string `json:"stream_arn,omitempty"`
	NoRotate   bool   `json:"no_rotate,omitempty"`
	MaxRetries *int   `json:"max_retries,omitempty"`

	BatchSize      int           `json:"batch_size,omitempty"`
	BatchSizeBytes int           `json:"batch_size_bytes,omitempty"`
	BatchTimeout   time.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.MaxRetries == nil {
		s.MaxRetries = new(int)
		*s.MaxRetries = 5
	}

	if s.BatchSize == 0 {
		s.BatchSize = 500
	}

	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = 5 << 20 // 5 MiB
	}

	if s.BatchTimeout == 0 {
		s.BatchTimeout = 20 * time.Second // 20s
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
