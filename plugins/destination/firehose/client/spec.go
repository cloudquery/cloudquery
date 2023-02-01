package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/filetypes"
)

type Spec struct {
	StreamARN  string `json:"stream_arn,omitempty"`
	NoRotate   bool   `json:"no_rotate,omitempty"`
	MaxRetries *int   `json:"max_retries,omitempty"`
	*filetypes.FileSpec
}

func (s *Spec) SetDefaults() {
	if s.MaxRetries == nil {
		s.MaxRetries = new(int)
		*s.MaxRetries = 5
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
