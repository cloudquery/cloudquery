package table_options

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/cloudquery/plugin-sdk/v3/caser"
	"github.com/jinzhu/copier"
)

type AccessanalyzerFindings struct {
	ListFindingOpts CustomAccessAnalyzerListFindingsInput `json:"list_findings,omitempty"`
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
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListFindingsInput)
}

func (c *AccessanalyzerFindings) validateListFindings() error {
	if aws.ToString(c.ListFindingOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in ListFindings")
	}
	if aws.ToString(c.ListFindingOpts.AnalyzerArn) != "" {
		return errors.New("invalid input: cannot set AnalyzerARN in ListFindings")
	}
	return nil
}

func (c *AccessanalyzerFindings) Validate() error {
	return c.validateListFindings()
}

func (c *AccessanalyzerFindings) ListFindings() (*accessanalyzer.ListFindingsInput, error) {
	var aaLFI accessanalyzer.ListFindingsInput
	if c == nil {
		return &aaLFI, nil
	}
	if err := c.validateListFindings(); err != nil {
		return &aaLFI, err
	}

	return &aaLFI, copier.Copy(&aaLFI, &c.ListFindingOpts)
}
