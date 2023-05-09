package client

import (
	"encoding/json"
	"testing"
	"time"
)

// TestTableOptionsUnmarshal tests that the TableOptions struct can be unmarshaled from JSON using
// snake_case keys.
func TestTableOptionsUnmarshal(t *testing.T) {
	opts := `{
	"aws_cloudtrail_events": {
		"lookup_events": {
			"start_time": "2020-01-01T00:00:00Z",
			"end_time": "2020-01-02T00:00:00Z",
			"event_category": "insight",
			"lookup_attributes": [
				{
					"attribute_key": "EventName",
					"attribute_value": "ConsoleLogin"
				}
			]
		}
	}
}`
	var tableOpts TableOptions
	if err := json.Unmarshal([]byte(opts), &tableOpts); err != nil {
		t.Fatal(err)
	}
	if tableOpts.CloudTrailEvents == nil {
		t.Fatal("CloudTrailEvents is nil")
	}
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	gotStart := tableOpts.CloudTrailEvents.LookupEventsOpts.StartTime
	if gotStart == nil || !gotStart.Equal(start) {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.StartTime = %s, want %s", tableOpts.CloudTrailEvents.LookupEventsOpts.StartTime, start)
	}
	end := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	gotEnd := tableOpts.CloudTrailEvents.LookupEventsOpts.EndTime
	if gotEnd == nil || !gotEnd.Equal(end) {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.EndTime = %s, want %s", tableOpts.CloudTrailEvents.LookupEventsOpts.EndTime, end)
	}
	if tableOpts.CloudTrailEvents.LookupEventsOpts.EventCategory != "insight" {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.EventCategory = %s, want insight", tableOpts.CloudTrailEvents.LookupEventsOpts.EventCategory)
	}
	if len(tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes) != 1 {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.LookupAttributes = %v, want 1 item", tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes)
	}
	if tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeKey != "EventName" {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeKey = %v, want EventName", tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeKey)
	}
	if *tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeValue != "ConsoleLogin" {
		t.Fatalf("CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeKey = %v, want EventName", *tableOpts.CloudTrailEvents.LookupEventsOpts.LookupAttributes[0].AttributeValue)
	}
}
