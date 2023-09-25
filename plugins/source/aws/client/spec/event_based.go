package spec

import "time"

type EventBasedSync struct {
	FullSync         *bool      `json:"full_sync,omitempty"`
	Account          Account    `json:"account"`
	KinesisStreamARN string     `json:"kinesis_stream_arn"`
	StartTime        *time.Time `json:"start_time,omitempty"`
}
