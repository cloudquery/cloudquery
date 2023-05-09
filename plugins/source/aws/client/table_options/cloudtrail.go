package table_options

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

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
	changeCaseForObject(m)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.LookupEventsInput)
}

type CloudtrailAPIs struct {
	LookupEventsOpts CustomLookupEventsOpts `json:"lookup_events,omitempty"`
}

func (c *CloudtrailAPIs) validateLookupEvents() error {
	if aws.ToString(c.LookupEventsOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in LookupEvents")
	}
	return nil
}

func (c *CloudtrailAPIs) Validate() error {
	return c.validateLookupEvents()
}
