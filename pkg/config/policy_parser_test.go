package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPolicyUnexpectedBlock = `policy "test_policy" {
  unknown {
    attr = "value"
  }
}`

const testPolicyMultipleConfigurationBlocks = `policy "test_policy" {
  configuration {
  }
  configuration {
  }
}`

const testPolicyQueries = `policy "test_policy" {
  query "first" {
    query = "query1"
  }
  query "second" {
    query = "query2"
    type = "automatic"
  }
  query "third" {
    query = "query3"
    type = "manual"
  }
}`

const testPolicyInvalidQueryType = `policy "test_policy" {
  query "first" {
    query = "query1"
    type = "invalid"
  }
}`

const testPolicy = `policy "aws-cis-v1.3.0" {
  description = "AWS CIS V1.3.0"
  configuration {
    provider "aws" {
      version = ">= 1.0"
    }
  }

  view "aws-cis-view" {
    description = "AWS CIS View"
    query "test-query-view" {
      query = "SELECT * FROM my.view"
    }
  }

  query "top-level-query" {
    description = "Top Level Query"
    query = "SELECT * FROM test"
    type = "manual"
  }

  policy "sub-policy-1" {
    description = "Sub Policy 1"
    query "sub-level-query" {
      query = "SELECT * from test.subquery"
    }
  }

  policy "sub-policy-2" {
    description = "Sub Policy 2"
    query "sub-level-query" {
      query = "SELECT * from test.subquery"
      type = "manual"
    }
  }
}`

func TestPolicyParser_LoadConfigFromSource(t *testing.T) {
	tests := []struct {
		name         string
		policyText   string
		expected     *PolicyWrapper
		wantErr      bool
		errHaystring string
	}{
		{
			"unexpected block within a policy",
			testPolicyUnexpectedBlock,
			nil,
			true,
			"Unsupported block type",
		},
		{
			"multiple configuration blocks",
			testPolicyMultipleConfigurationBlocks,
			nil,
			true,
			"Duplicate block",
		},
		{
			"queries with or without explicit type",
			testPolicyQueries,
			&PolicyWrapper{Policies: []*Policy{{
				Name: "test_policy",
				Queries: []*Query{
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
			}}},
			false,
			"",
		},
		{
			"query with invalid type",
			testPolicyInvalidQueryType,
			nil,
			true,
			"Invalid query type",
		},
		{
			"complex policy",
			testPolicy,
			&PolicyWrapper{Policies: []*Policy{{
				Name:        "aws-cis-v1.3.0",
				Description: "AWS CIS V1.3.0",
				Config: &Configuration{
					Providers: []*PolicyProvider{{
						Type:    "aws",
						Version: ">= 1.0",
					}},
				},
				Views: []*View{{
					Name:        "aws-cis-view",
					Description: "AWS CIS View",
					Query: &Query{
						Name:  "test-query-view",
						Query: "SELECT * FROM my.view",
					},
				}},
				Queries: []*Query{{
					Name:        "top-level-query",
					Description: "Top Level Query",
					Type:        ManualQuery,
					Query:       "SELECT * FROM test",
				}},
				Policies: []*Policy{
					{
						Name:        "sub-policy-1",
						Description: "Sub Policy 1",
						Queries: []*Query{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  AutomaticQuery,
						}},
					},
					{
						Name:        "sub-policy-2",
						Description: "Sub Policy 2",
						Queries: []*Query{{
							Name:  "sub-level-query",
							Query: "SELECT * from test.subquery",
							Type:  ManualQuery,
						}},
					},
				},
			}}},
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(nil)
			policiesRaw, diags := p.loadFromSource("policy.hcl", []byte(tt.policyText), SourceHCL)
			if diags != nil && diags.HasErrors() {
				t.Fatal(diags.Errs())
			}
			policiesWrapper, diags := p.DecodePolicies(policiesRaw, diags, "")
			if tt.wantErr != diags.HasErrors() {
				t.Errorf("want errors is %v, but have %v, error details: %s", tt.wantErr, diags.HasErrors(), diags.Error())
			}
			if tt.errHaystring != "" {
				assert.Contains(t, diags.Error(), tt.errHaystring)
			}
			assert.Equal(t, tt.expected, policiesWrapper)
		})
	}
}
