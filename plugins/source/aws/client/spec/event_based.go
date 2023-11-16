package spec

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

// Event-based sync configuration.
// This feature is available only in premium version of the plugin.
type EventBasedSync struct {
	// Whether the full sync will be performed for the tables prior to engaging the event-based sync mode.
	FullSync *bool `json:"full_sync,omitempty" jsonschema:"default=true"`

	// Account spec to configure sync.
	Account Account `json:"account"`

	// Amazon Kinesis stream ARN to subscribe to.
	KinesisStreamARN string `json:"kinesis_stream_arn" jsonschema:"required,pattern=^arn(:[^:\n]*){5}([:/].*)?$"`

	// The timestamp of the data record from which to start reading.
	StartTime *time.Time `json:"start_time,omitempty" jsonschema:"default=now"`
}

func (e *EventBasedSync) Validate() error {
	if !arn.IsARN(e.KinesisStreamARN) {
		return fmt.Errorf("kinesis_stream_arn %q is not a valid ARN", e.KinesisStreamARN)
	}
	return nil
}
