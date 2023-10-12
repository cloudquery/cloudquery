package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type AccessAnalyzerFindings struct {
	ListFindingsOpts []CustomAccessAnalyzerListFindingsInput `json:"list_findings,omitempty"`
}

type CustomAccessAnalyzerListFindingsInput struct {
	accessanalyzer.ListFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomAccessAnalyzerListFindingsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomAccessAnalyzerListFindingsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	skipFields := []string{"filter"}
	changeCaseForObject(m, csr.ToPascal, skipFields...)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListFindingsInput)
}

func (CustomAccessAnalyzerListFindingsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	// The following properties are prohibited in spec
	sc.Properties.Delete("NextToken")
	sc.Properties.Delete("AnalyzerArn")
}

func (s *AccessAnalyzerFindings) validateListFindings() error {
	for _, opt := range s.ListFindingsOpts {
		if opt.NextToken != nil {
			return errors.New("invalid input: cannot set NextToken in ListFindings")
		}
		if opt.AnalyzerArn != nil {
			return errors.New("invalid input: cannot set AnalyzerArn in ListFindings")
		}
	}
	return nil
}

func (s *AccessAnalyzerFindings) sanitized() *AccessAnalyzerFindings {
	var result AccessAnalyzerFindings
	if s != nil {
		result = *s
	}

	if len(result.ListFindingsOpts) == 0 {
		result.ListFindingsOpts = []CustomAccessAnalyzerListFindingsInput{{ListFindingsInput: accessanalyzer.ListFindingsInput{}}}
	}
	return &result
}

func (s *AccessAnalyzerFindings) Validate() error {
	return s.sanitized().validateListFindings()
}

func (s *AccessAnalyzerFindings) Filters() []CustomAccessAnalyzerListFindingsInput {
	if s != nil && s.ListFindingsOpts != nil {
		return s.ListFindingsOpts
	}
	return s.sanitized().ListFindingsOpts
}
