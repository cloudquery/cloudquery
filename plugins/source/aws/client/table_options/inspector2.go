package table_options

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
)

type Inspector2APIs struct {
	ListFindingsOpts CustomInspector2ListFindingsInput `json:"list_findings,omitempty"`
}

type CustomInspector2ListFindingsInput struct {
	inspector2.ListFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomLookupEventsOpts type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomInspector2ListFindingsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	changeCaseForObject(m)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListFindingsInput)
}

func (c *Inspector2APIs) validateListFindings() error {
	if aws.ToString(c.ListFindingsOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	return nil
}

func (c *Inspector2APIs) Validate() error {
	return c.validateListFindings()
}
