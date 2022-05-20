package policy

import (
	"fmt"
	"strings"
	"testing"
)

func TestPolicyParser_LoadConfigFromSource(t *testing.T) {
	tests := []struct {
		name                     string
		policyHCL                string
		expected                 *Policy
		expectedError            bool
		expectedErrorStr         string
		expectedValidationErrors []string
	}{
		{
			name:                     "test empty policy hcl",
			policyHCL:                "",
			expected:                 &Policy{},
			expectedValidationErrors: []string{"name is required"},
		},
		{
			name: "test policy with source",
			policyHCL: `
name: "source_policy"
source: "./path/to/file"
`,
			expected: &Policy{
				Name:   "source_policy",
				Source: "./path/to/file",
			},
		},
		{
			name: "illegal test policy with source",
			policyHCL: `
name: "source_policy"
source: "./path/to/file"
check:
  - name: "sub-level-query"
    query = "SELECT * from test.subquery"
    type = "manual"
`,
			expectedError:    true,
			expectedErrorStr: "Found source with blocks; There must be one of the following: Policy source attribute or blocks",
		},
		{
			name: "unexpected block within a policy",
			policyHCL: `
name: "test_policy":
unknown:
  attr: "value"`,
			expectedError:    true,
			expectedErrorStr: "Unsupported block type",
		},
		{
			name: "multiple configuration blocks",
			policyHCL: `
policy: "test_policy"
  configuration:
  configuration:`,
			expected: &Policy{
				Name:   "test_policy",
				Config: &Configuration{},
			},
			expectedError:    true,
			expectedErrorStr: "Duplicate block",
		},
		{
			name: "queries with or without explicit type",
			policyHCL: `
title: "test_policy"
checks:
  - name: "first"
    query: "query1"
  - name: "second"
      query: "query2"
    type: "automatic"
  - name: "third"
    query: "query3"
    type: "manual"
`,
			expected: &Policy{
				Name: "test_policy",
				Checks: []*Check{
					{
						Name:  "first",
						Query: "query1",
						Type:  AutomaticQuery,
					},
					{
						Name:  "second",
						Query: "query2",
						Type:  AutomaticQuery,
					},
					{
						Name:  "third",
						Query: "query3",
						Type:  ManualQuery,
					},
				},
			},
		},
		{
			name: "query with invalid type",
			policyHCL: `
title: "test_policy"
  checks:
    - name: "first"
      query: "query1"
      type: "invalid"
`,
			expected: &Policy{
				Name: "test_policy",
				Checks: []*Check{
					{
						Name:  "first",
						Type:  "invalid",
						Query: "query1",
					},
				},
			},
			expectedError:    true,
			expectedErrorStr: "Invalid query type",
		},
		{
			name: "complex policy",
			policyHCL: `
name: "aws-cis-v1.3.0"
title = "AWS CIS V1.3.0"
doc = "some doc info"
configuration:
  providers:
    - name: "aws"
      version:  ">= 1.0"

  views:
    - name: "aws-cis-view"
      title: "AWS CIS View"
      query: "SELECT * FROM my.view"

  check:
    - name: "top-level-query"
      title: "Top Level Check"
      query: "SELECT * FROM test"
      type: "manual"

  policies:
    - name: "sub-policy-1"
      title: "Sub Policy 1"
      checks:
        - name: "sub-level-query"
          query: "SELECT * from test.subquery"

    - name: "sub-policy-2":
      title: "Sub Policy 2"
      checks:
        - name: "sub-level-query"
          query: "SELECT * from test.subquery"
          type: "manual"`,
			expected: &Policy{
				Name:  "aws-cis-v1.3.0",
				Title: "AWS CIS V1.3.0",
				Doc:   "some doc info",
				Config: &Configuration{
					Providers: []*Provider{{
						Type:    "aws",
						Version: ">= 1.0",
					}},
				},
				Views: []*View{{
					Name:  "aws-cis-view",
					Title: "AWS CIS View",
					Query: "SELECT * FROM my.view",
				}},
				Checks: []*Check{{
					Name:  "top-level-query",
					Title: "Top Level Check",
					Type:  ManualQuery,
					Query: "SELECT * FROM test",
				}},
				Policies: []*Policy{
					{
						Name:  "sub-policy-1",
						Title: "Sub Policy 1",
						Checks: []*Check{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  AutomaticQuery,
						}},
					},
					{
						Name:  "sub-policy-2",
						Title: "Sub Policy 2",
						Checks: []*Check{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  ManualQuery,
						}},
					},
				},
			},
		},
		{
			name: "policy with slash in the label",
			policyHCL: `
name: "po/licy"
  checks:
    - name: "sub-level-query"
      query: "SELECT * from test.subquery"
      type: "manual"`,
			expected: &Policy{
				Name: "po/licy",
				Checks: []*Check{
					{
						Name:  "sub-level-query",
						Query: "SELECT * from test.subquery",
						Type:  "manual",
					},
				},
			},
			expectedError:    true,
			expectedErrorStr: "Slash character in policy label",
		},
		{
			name: "check with slash in the label",
			policyHCL: `
      name: "policy"
        checks:
          - name: "sub/level-query"
            query: "SELECT * from test.subquery"
            type: "manual"
  `,
			expected: &Policy{
				Name: "policy",
				Checks: []*Check{
					{
						Name:  "sub/level-query",
						Query: "SELECT * from test.subquery",
						Type:  "manual",
					},
				},
			},
			expectedError:    true,
			expectedErrorStr: "Slash character in check label",
		},
		{
			name: "test-policy-identifiers",
			policyHCL: `
title: "test_policy"
identifiers: ["id"]
checks:
  - name: "1"
    query: "select 1 as id, 'test' as cq_reason"
policies:
  - name: child_policy
    check:
      - name: child_check
        query: "select 1 as id"
        reason: "test"`,
			expected: &Policy{
				Name:        "test_policy",
				Identifiers: []string{"id"},
				Checks: []*Check{
					{
						Name:  "1",
						Query: "select 1 as id, 'test' as cq_reason",
						Type:  AutomaticQuery,
					},
				},
				Policies: Policies{
					{
						Name: "child_policy",
						Checks: []*Check{
							{
								Name:   "1",
								Query:  "select 1 as id",
								Type:   AutomaticQuery,
								Reason: "test",
							},
						},
						Identifiers: []string{"id"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, result, err := UnmarshalPolicy([]byte(tt.policyHCL))
			if err != nil {
				t.Fatal(err)
			}
			if !result.Valid() {
				if len(result.Errors()) != len(tt.expectedValidationErrors) {
					t.Fatalf("Expected %d validation errors, got %d", len(tt.expectedValidationErrors), len(result.Errors()))
				}
				for i, err := range result.Errors() {
					if strings.Compare(err.Description(), tt.expectedValidationErrors[i]) != 0 {
						t.Fatalf("Expected validation error %s, got %s", tt.expectedValidationErrors[i], err.Description())
					}
				}
			}
			fmt.Println(p)

			// if diff := cmp.Diff(p, tt.expected); diff != "" {
			// 	t.Errorf("Config mismatch (-want +got):\n%s", diff)
			// }

		})
	}
}
