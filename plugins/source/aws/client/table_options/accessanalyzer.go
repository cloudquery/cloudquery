package table_options

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options/inputs/accessanalyzer_input"
	"github.com/jinzhu/copier"
)

type AaFindings struct {
	ListFindingOpts accessanalyzer_input.ListFindingsInput `json:"list_findings,omitempty"`
}

func (c *AaFindings) validateListFindings() error {
	if aws.ToString(c.ListFindingOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	if aws.ToString(c.ListFindingOpts.AnalyzerArn) != "" {
		return errors.New("invalid input: cannot set AnalyzerARN in ListFindings")
	}
	return nil
}

func (c *AaFindings) ListFindings() (*accessanalyzer.ListFindingsInput, error) {
	var aaLFI accessanalyzer.ListFindingsInput
	if c == nil {
		return &aaLFI, nil
	}
	if err := c.validateListFindings(); err != nil {
		return &aaLFI, err
	}

	return &aaLFI, copier.Copy(&aaLFI, &c.ListFindingOpts)
}
