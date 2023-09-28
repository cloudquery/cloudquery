package spec

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type EventBasedSync struct {
	FullSync         *bool      `json:"full_sync,omitempty" jsonschema:"default=true"`
	Account          Account    `json:"account"`
	KinesisStreamARN string     `json:"kinesis_stream_arn" jsonschema:"required,pattern=^arn(:[^:\n]*){5\\,}([:/].*)?$"`
	StartTime        *time.Time `json:"start_time,omitempty" jsonschema:"default=now"`
}

func (e *EventBasedSync) Validate() error {
	if !arn.IsARN(e.KinesisStreamARN) {
		return fmt.Errorf("kinesis_stream_arn %q is not a valid ARN", e.KinesisStreamARN)
	}
	return nil
}
