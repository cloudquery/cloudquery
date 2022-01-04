package policy

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
)

const (
	testPolicyUnexpectedBlock = `policy "test_policy" {
 unknown {
   attr = "value"
 }
}`
	testPolicyMultipleConfigurationBlocks = `policy "test_policy" {
 configuration {
 }
 configuration {
 }
}`
	testPolicyQueries = `policy "test_policy" {
 check "first" {
   query = "query1"
 }
 check "second" {
   query = "query2"
   type = "automatic"
 }
 check "third" {
   query = "query3"
   type = "manual"
 }
}`
	testPolicyInvalidQueryType = `policy "test_policy" {
 check "first" {
   query = "query1"
   type = "invalid"
 }
}`

	testPolicy = `policy "aws-cis-v1.3.0" {
 description = "AWS CIS V1.3.0"
 doc = "some doc info"
 configuration {
   provider "aws" {
     version = ">= 1.0"
   }
 }

 view "aws-cis-view" {
   description = "AWS CIS View"
   query = "SELECT * FROM my.view"
 }

 check "top-level-query" {
   description = "Top Level Check"
   query = "SELECT * FROM test"
   type = "manual"
 }

 policy "sub-policy-1" {
   description = "Sub Policy 1"
   check "sub-level-query" {
     query = "SELECT * from test.subquery"
   }
 }

 policy "sub-policy-2" {
   description = "Sub Policy 2"
   check "sub-level-query" {
     query = "SELECT * from test.subquery"
     type = "manual"
   }
 }
}`
	testSinglePolicyBlock = `
		policy "source_policy" {
			source = "./path/to/file"
		}
		policy "source_policy" {
			source = "./path/to/file"
		}`

	testPolicySource = `
		policy "source_policy" {
			source = "./path/to/file"
		}`

	testInvalidPolicySource = `
		policy "source_policy" {
			source = "./path/to/file"
			check "sub-level-query" {
				query = "SELECT * from test.subquery"
				type = "manual"
			}
		}`
)

func TestPolicyParser_LoadConfigFromSource(t *testing.T) {
	tests := []struct {
		name      string
		policyHCL string
		expected  *Policy
		wantErr   bool
		errString string
	}{
		{
			name:      "test policy with source",
			policyHCL: testPolicySource,
			expected: &Policy{
				Name:   "source_policy",
				Source: "./path/to/file",
			},
		},
		{
			name:      "single root policy block",
			policyHCL: testSinglePolicyBlock,
			wantErr:   true,
			errString: "Only single root policy block allowed; Only a single policy block is allowed in root level policy",
		},
		{
			name:      "illegal test policy with source",
			policyHCL: testInvalidPolicySource,
			wantErr:   true,
			errString: "Found source with blocks; There must be one of the following: Policy source attribute or blocks",
		},
		{
			name:      "unexpected block within a policy",
			policyHCL: testPolicyUnexpectedBlock,
			wantErr:   true,
			errString: "Unsupported block type",
		},
		{
			name:      "multiple configuration blocks",
			policyHCL: testPolicyMultipleConfigurationBlocks,
			wantErr:   true,
			errString: "Duplicate block",
		},
		{
			name:      "queries with or without explicit type",
			policyHCL: testPolicyQueries,
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
			name:      "query with invalid type",
			policyHCL: testPolicyInvalidQueryType,
			wantErr:   true,
			errString: "Invalid query type",
		},
		{
			name:      "complex policy",
			policyHCL: testPolicy,
			expected: &Policy{
				Name:        "aws-cis-v1.3.0",
				Description: "AWS CIS V1.3.0",
				Doc:         "some doc info",
				Config: &Configuration{
					Providers: []*Provider{{
						Type:    "aws",
						Version: ">= 1.0",
					}},
				},
				Views: []*View{{
					Name:        "aws-cis-view",
					Description: "AWS CIS View",
					Query:       "SELECT * FROM my.view",
				}},
				Checks: []*Check{{
					Name:        "top-level-query",
					Description: "Top Level Check",
					Type:        ManualQuery,
					Query:       "SELECT * FROM test",
				}},
				Policies: []*Policy{
					{
						Name:        "sub-policy-1",
						Description: "Sub Policy 1",
						Checks: []*Check{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  AutomaticQuery,
						}},
					},
					{
						Name:        "sub-policy-2",
						Description: "Sub Policy 2",
						Checks: []*Check{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  ManualQuery,
						}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, diags := hclsyntax.ParseConfig([]byte(tt.policyHCL), t.Name(), hcl.Pos{Byte: 0, Line: 1, Column: 1})
			if diags != nil && diags.HasErrors() {
				t.Fatal(diags.Errs())
			}
			policiesWrapper, diags := DecodePolicy(f.Body, diags, "")
			if tt.wantErr != diags.HasErrors() {
				t.Errorf("want errors is %v, but have %v, error details: %s", tt.wantErr, diags.HasErrors(), diags.Error())
			}
			if tt.errString != "" {
				assert.Contains(t, diags.Error(), tt.errString)
			}
			assert.Equal(t, tt.expected, policiesWrapper)
		})
	}
}
