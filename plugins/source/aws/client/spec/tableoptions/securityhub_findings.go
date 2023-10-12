package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type SecurityHubFindings struct {
	GetFindingsOpts []CustomSecurityHubGetFindingsInput `json:"get_findings,omitempty"`
}

type CustomSecurityHubGetFindingsInput struct {
	securityhub.GetFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomSecurityHubGetFindingsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomSecurityHubGetFindingsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &s.GetFindingsInput)
}

func (CustomSecurityHubGetFindingsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	// The following properties are prohibited in spec
	sc.Properties.Delete("NextToken")
	// The following properties have additional constraints
	propertyMaxResults := sc.Properties.Value("MaxResults")
	if len(propertyMaxResults.OneOf) == 2 {
		propertyMaxResults = propertyMaxResults.OneOf[0] // 0 = value, 1 = null
	}
	propertyMaxResults.Minimum = json.Number("1")
	propertyMaxResults.Maximum = json.Number("100")
	propertyMaxResults.Default = 100
}

func (s *SecurityHubFindings) validateGetFindings() error {
	for _, opt := range s.GetFindingsOpts {
		if opt.NextToken != nil {
			return errors.New("invalid input: cannot set NextToken in GetFindings")
		}
		if opt.MaxResults < 1 || opt.MaxResults > 100 {
			return errors.New("invalid range: MaxResults must be within range [1-100]")
		}
	}
	return nil
}

func (s *SecurityHubFindings) sanitized() *SecurityHubFindings {
	var result SecurityHubFindings
	if s != nil {
		result = *s
	}

	if len(result.GetFindingsOpts) == 0 {
		result.GetFindingsOpts = []CustomSecurityHubGetFindingsInput{{GetFindingsInput: securityhub.GetFindingsInput{}}}
	}
	for i, opt := range result.GetFindingsOpts {
		if opt.MaxResults == 0 {
			result.GetFindingsOpts[i].MaxResults = 100
		}
	}
	return &result
}

func (s *SecurityHubFindings) Validate() error {
	return s.sanitized().validateGetFindings()
}

func (s *SecurityHubFindings) Filters() []CustomSecurityHubGetFindingsInput {
	if s != nil && s.GetFindingsOpts != nil {
		return s.GetFindingsOpts
	}
	return s.sanitized().GetFindingsOpts
}
