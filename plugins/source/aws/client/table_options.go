package client

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailTypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
)

type TableOptions struct {
	CloudTrailEvents   *ctAPIs         `json:"aws_cloudtrail_events,omitempty"`
	Inspector2Findings *inspector2APIs `json:"aws_inspector2_findings,omitempty"`
}

type ctAPIs struct {
	LookupEventsOpts LookupEventsOptions `json:"LookupEvents,omitempty"`
}

type LookupEventsOptions struct {
	EndTime          *time.Time
	EventCategory    cloudtrailTypes.EventCategory
	LookupAttributes []cloudtrailTypes.LookupAttribute
	StartTime        *time.Time
}

func (c *ctAPIs) LookupEvents() *cloudtrail.LookupEventsInput {
	if c == nil {
		return &cloudtrail.LookupEventsInput{}
	}

	return &cloudtrail.LookupEventsInput{
		EndTime:          c.LookupEventsOpts.EndTime,
		EventCategory:    c.LookupEventsOpts.EventCategory,
		LookupAttributes: c.LookupEventsOpts.LookupAttributes,
		StartTime:        c.LookupEventsOpts.StartTime,
	}
}

type inspector2APIs struct {
	ListFindingOpts inspector2ListFindingsOptions `json:"ListFindings,omitempty"`
}

type inspector2ListFindingsOptions struct {
	FilterCriteria *types.FilterCriteria
	SortCriteria   *types.SortCriteria
}

func (c *inspector2APIs) ListFindings() *inspector2.ListFindingsInput {
	if c == nil {
		return &inspector2.ListFindingsInput{}
	}
	return &inspector2.ListFindingsInput{
		FilterCriteria: c.ListFindingOpts.FilterCriteria,
	}
}
