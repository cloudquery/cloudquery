package table_options

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/table_options/inputs/cloudtrail_input"
	"github.com/jinzhu/copier"
)

type CtAPIs struct {
	LookupEventsOpts cloudtrail_input.LookupEventsInput `json:"LookupEvents,omitempty"`
}

func (c *CtAPIs) validateLookupEvents() error {
	if aws.ToString(c.LookupEventsOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in LookupEvents")
	}
	return nil
}

func (c *CtAPIs) LookupEvents() (*cloudtrail.LookupEventsInput, error) {
	var lookupEventsInput cloudtrail.LookupEventsInput
	if c == nil {
		return &lookupEventsInput, nil
	}
	// validate input
	if err := c.validateLookupEvents(); err != nil {
		return &lookupEventsInput, err
	}

	// copy input to AWS type

	return &lookupEventsInput, copier.Copy(&lookupEventsInput, &c.LookupEventsOpts)
}
