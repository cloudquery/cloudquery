package client

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
)

type TableOptions struct {
	CloudTrailEvents       *ctAPIs         `json:"aws_cloudtrail_events,omitempty"`
	AccessAnalyzerFindings *aaFindings     `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	Inspector2Findings     *inspector2APIs `json:"aws_inspector2_findings,omitempty"`
}

type aaFindings struct {
	ListFindingOpts accessanalyzer.ListFindingsInput `json:"ListFindings,omitempty"`
}

func (c *aaFindings) ListFindings() (*accessanalyzer.ListFindingsInput, error) {
	if c == nil {
		return &accessanalyzer.ListFindingsInput{}, nil
	}

	if aws.ToString(c.ListFindingOpts.NextToken) != "" {
		return &accessanalyzer.ListFindingsInput{}, errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	if aws.ToString(c.ListFindingOpts.AnalyzerArn) != "" {
		return &accessanalyzer.ListFindingsInput{}, errors.New("invalid input: cannot set AnalyzerARN in ListFindings")
	}
	return &c.ListFindingOpts, nil
}

type ctAPIs struct {
	LookupEventsOpts cloudtrail.LookupEventsInput `json:"LookupEvents,omitempty"`
}

func (c *ctAPIs) LookupEvents() (*cloudtrail.LookupEventsInput, error) {
	if c == nil {
		return &cloudtrail.LookupEventsInput{}, nil
	}
	if aws.ToString(c.LookupEventsOpts.NextToken) != "" {
		return &cloudtrail.LookupEventsInput{}, errors.New("invalid input: cannot set NextToken in LookupEvents")
	}
	return &c.LookupEventsOpts, nil
}

type inspector2APIs struct {
	ListFindingOpts inspector2.ListFindingsInput `json:"ListFindings,omitempty"`
}

func (c *inspector2APIs) ListFindings() (*inspector2.ListFindingsInput, error) {
	if c == nil {
		return &inspector2.ListFindingsInput{}, nil
	}

	if aws.ToString(c.ListFindingOpts.NextToken) != "" {
		return &inspector2.ListFindingsInput{}, errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	return &c.ListFindingOpts, nil
}
