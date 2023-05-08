package table_options

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options/inputs/inspector2_input"
	"github.com/jinzhu/copier"
)

type Inspector2APIs struct {
	ListFindingOpts inspector2_input.ListFindingsInput `json:"list_findings,omitempty"`
}

func (c *Inspector2APIs) validateListFindings() error {
	if aws.ToString(c.ListFindingOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	return nil
}

func (c *Inspector2APIs) ListFindings() (*inspector2.ListFindingsInput, error) {
	var inspector2LFI inspector2.ListFindingsInput
	if c == nil {
		return &inspector2LFI, nil
	}
	// validate input
	if err := c.validateListFindings(); err != nil {
		return &inspector2.ListFindingsInput{}, err
	}

	// copy input to AWS type
	return &inspector2LFI, copier.Copy(&inspector2LFI, &c.ListFindingOpts)
}
