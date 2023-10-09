package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type CloudtrailEvents struct {
	LookupEventsOpts []CustomCloudtrailLookupEventsInput `json:"lookup_events,omitempty"`
}

type CustomCloudtrailLookupEventsInput struct {
	cloudtrail.LookupEventsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomCloudtrailLookupEventsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomCloudtrailLookupEventsInput) UnmarshalJSON(data []byte) error {
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

// JSONSchemaExtend is required to remove `NextToken`.
func (CustomCloudtrailLookupEventsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("NextToken")
}

func (c *CloudtrailEvents) validateLookupEvents() error {
	for _, opt := range c.LookupEventsOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in LookupEvents")
		}
	}
	return nil
}

func (c *CloudtrailEvents) Validate() error {
	return c.validateLookupEvents()
}
