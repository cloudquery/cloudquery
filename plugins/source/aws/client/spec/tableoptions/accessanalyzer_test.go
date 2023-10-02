package tableoptions

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudquery/plugin-sdk/faker"
)

func TestAAListFindings(t *testing.T) {
	u := CustomAccessAnalyzerListFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))

	api := AccessanalyzerFindings{
		ListFindingOpts: []CustomAccessAnalyzerListFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in ListFindings")
	api.ListFindingOpts[0].NextToken = nil

	err = api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set AnalyzerARN in ListFindings")
	api.ListFindingOpts[0].AnalyzerArn = nil

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	err = api.Validate()
	assert.Nil(t, err)
}

func TestAccessAnalyzerJSONSchema(t *testing.T) {
	testJSONSchema(t, []jsonSchemaTestCase{
		{
			name: "empty",
			spec: `{"aws_accessanalyzer_analyzer_findings":{}}`,
		},
		{
			name: "proper",
			spec: func() string {
				var findings AccessanalyzerFindings
				require.NoError(t, faker.FakeObject(&findings))
				require.Len(t, findings.ListFindingOpts, 1)

				// remove prohibited fields
				findings.ListFindingOpts[0].AnalyzerArn = nil
				findings.ListFindingOpts[0].NextToken = nil

				data, err := json.MarshalIndent(TableOptions{AccessAnalyzerFindings: &findings}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				result = strings.Replace(result, "\"AnalyzerArn\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "AnalyzerArn is present",
			err:  true,
			spec: func() string {
				var findings AccessanalyzerFindings
				require.NoError(t, faker.FakeObject(&findings))
				require.Len(t, findings.ListFindingOpts, 1)

				// remove prohibited fields
				findings.ListFindingOpts[0].NextToken = nil

				data, err := json.MarshalIndent(TableOptions{AccessAnalyzerFindings: &findings}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				return result
			}(),
		},
		{
			name: "NextToken is present",
			err:  true,
			spec: func() string {
				var findings AccessanalyzerFindings
				require.NoError(t, faker.FakeObject(&findings))
				require.Len(t, findings.ListFindingOpts, 1)

				// remove prohibited fields
				findings.ListFindingOpts[0].AnalyzerArn = nil

				data, err := json.MarshalIndent(TableOptions{AccessAnalyzerFindings: &findings}, "", "  ")
				require.NoError(t, err)
				result := string(data)
				result = strings.Replace(result, "\"NextToken\": null,\n", ``, 1)
				return result
			}(),
		},
	})
}
