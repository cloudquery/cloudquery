package tableoptions

import (
	"encoding/json"
	"errors"

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
func (s *CustomCloudtrailLookupEventsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &s.LookupEventsInput)
}

func (CustomCloudtrailLookupEventsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	// The following properties are prohibited in spec
	sc.Properties.Delete("NextToken")
}

func (s *CloudtrailEvents) validateLookupEvents() error {
	for _, opt := range s.LookupEventsOpts {
		if opt.NextToken != nil {
			return errors.New("invalid input: cannot set NextToken in LookupEvents")
		}
	}
	return nil
}

func (s *CloudtrailEvents) sanitized() *CloudtrailEvents {
	var result CloudtrailEvents
	if s != nil {
		result = *s
	}

	if len(result.LookupEventsOpts) == 0 {
		result.LookupEventsOpts = []CustomCloudtrailLookupEventsInput{{LookupEventsInput: cloudtrail.LookupEventsInput{}}}
	}
	return &result
}

func (s *CloudtrailEvents) Validate() error {
	return s.sanitized().validateLookupEvents()
}

func (s *CloudtrailEvents) Filters() []CustomCloudtrailLookupEventsInput {
	if s != nil && s.LookupEventsOpts != nil {
		return s.LookupEventsOpts
	}
	return s.sanitized().LookupEventsOpts
}
