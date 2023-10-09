package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type AccessAnalyzerFindings struct {
	ListFindingOpts []CustomAccessAnalyzerListFindingsInput `json:"list_findings,omitempty"`
}

type CustomAccessAnalyzerListFindingsInput struct {
	accessanalyzer.ListFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomLookupEventsOpts type.
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

// JSONSchemaExtend is required to remove `AnalyzerArn` & `NextToken`.
func (CustomAccessAnalyzerListFindingsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("AnalyzerArn")
	sc.Properties.Delete("NextToken")
}

func (c *AccessAnalyzerFindings) validateListFindings() error {
	for _, opt := range c.ListFindingOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in ListFindings")
		}
		if aws.ToString(opt.AnalyzerArn) != "" {
			return errors.New("invalid input: cannot set AnalyzerARN in ListFindings")
		}
	}
	return nil
}

func (c *AccessAnalyzerFindings) Validate() error {
	return c.validateListFindings()
}
