package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/filetypes/v2"
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
	if s.FileSpec == nil {
		s.FileSpec = &filetypes.FileSpec{}
	}
	s.FileSpec.SetDefaults()
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
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}
	return s.FileSpec.Validate()
}
