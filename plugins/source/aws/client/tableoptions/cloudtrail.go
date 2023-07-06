package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/cloudquery/plugin-sdk/v4/caser"
)

type CloudtrailAPIs struct {
	LookupEventsOpts []CustomLookupEventsOpts `json:"lookup_events,omitempty"`
}

type CustomLookupEventsOpts struct {
	cloudtrail.LookupEventsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomLookupEventsOpts type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomLookupEventsOpts) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.LookupEventsInput)
}

func (c *CloudtrailAPIs) validateLookupEvents() error {
	for _, opt := range c.LookupEventsOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in LookupEvents")
		}
	}
	return nil
}

func (c *CloudtrailAPIs) Validate() error {
	return c.validateLookupEvents()
}
