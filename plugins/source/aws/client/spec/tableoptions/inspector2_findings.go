package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type Inspector2Findings struct {
	ListFindingsOpts []CustomInspector2ListFindingsInput `json:"list_findings,omitempty"`
}

type CustomInspector2ListFindingsInput struct {
	inspector2.ListFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomInspector2ListFindingsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomInspector2ListFindingsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &s.ListFindingsInput)
}

func (CustomInspector2ListFindingsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	// The following properties are prohibited in spec
	sc.Properties.Delete("NextToken")
}

func (s *Inspector2Findings) validateListFindings() error {
	for _, opt := range s.ListFindingsOpts {
		if opt.NextToken != nil {
			return errors.New("invalid input: cannot set NextToken in ListFindings")
		}
	}
	return nil
}

func (s *Inspector2Findings) sanitized() *Inspector2Findings {
	var result Inspector2Findings
	if s != nil {
		result = *s
	}

	if len(result.ListFindingsOpts) == 0 {
		result.ListFindingsOpts = []CustomInspector2ListFindingsInput{{ListFindingsInput: inspector2.ListFindingsInput{}}}
	}
	return &result
}

func (s *Inspector2Findings) Validate() error {
	return s.sanitized().validateListFindings()
}

func (s *Inspector2Findings) Filters() []CustomInspector2ListFindingsInput {
	if s != nil && s.ListFindingsOpts != nil {
		return s.ListFindingsOpts
	}
	return s.sanitized().ListFindingsOpts
}
