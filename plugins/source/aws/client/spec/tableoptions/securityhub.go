package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type SecurityHubAPIs struct {
	GetFindingsOpts []CustomGetFindingsOpts `json:"get_findings,omitempty"`
}

type CustomGetFindingsOpts struct {
	securityhub.GetFindingsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomGetFindingsOpts type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomGetFindingsOpts) UnmarshalJSON(data []byte) error {
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

// JSONSchemaExtend is required to remove `NextToken` as well as add min & max for `MaxResults`.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
func (CustomGetFindingsOpts) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("NextToken")

	maxResults := sc.Properties.Value("MaxResults")
	maxResults.Minimum = json.Number("1")
	maxResults.Maximum = json.Number("100")
}

func (s *SecurityHubAPIs) validateGetFindingEvent() error {
	for _, opt := range s.GetFindingsOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in GetFindings")
		}

		// As per https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html#API_GetFindings_RequestSyntax
		if opt.MaxResults < 1 || opt.MaxResults > 100 {
			return errors.New("invalid range: MaxResults must be within range [1-100]")
		}
	}
	return nil
}

func (s *SecurityHubAPIs) SetDefaults() {
	for i := 0; i < len(s.GetFindingsOpts); i++ {
		if s.GetFindingsOpts[i].MaxResults == 0 {
			s.GetFindingsOpts[i].MaxResults = 100
		}
	}
}

func (s *SecurityHubAPIs) Validate() error {
	return s.validateGetFindingEvent()
}
