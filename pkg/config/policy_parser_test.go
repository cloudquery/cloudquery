package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
    }
  }
}`

func TestPolicyParser_LoadConfigFromSource(t *testing.T) {
	p := NewParser(nil)
	policiesRaw, diags := p.loadFromSource("policy.hcl", []byte(testPolicy), SourceHCL)
	if diags != nil && diags.HasErrors() {
		t.Fatal(diags.Errs())
	}
	policiesWrapper, diags := p.DecodePolicies(policiesRaw, diags)

	// Only one module should be defined for this test case
	assert.Equal(t, 1, len(policiesWrapper.Policies))

	// Define expected struct
	exp := &Policy{
		Name:        "aws-cis-v1.3.0",
		Description: "AWS CIS V1.3.0",
		Config: &Configuration{
			Providers: []*PolicyProvider{
				{
					Type:    "aws",
					Version: ">= 1.0",
				},
			},
		},
		Policies: []*Policy{
			{
				Name:        "sub-policy-1",
				Description: "Sub Policy 1",
				Queries: []*Query{
					{
						Name:         "sub-level-query",
						ExpectOutput: false,
						Query:        "SELECT * from test.subquery",
					},
				},
			},
			{
				Name:        "sub-policy-2",
				Description: "Sub Policy 2",
				Queries: []*Query{
					{
						Name:         "sub-level-query",
						ExpectOutput: false,
						Query:        "SELECT * from test.subquery",
					},
				},
			},
		},
		Queries: []*Query{
			{
				Name:         "top-level-query",
				Description:  "Top Level Query",
				ExpectOutput: false,
				Query:        "SELECT * FROM test",
			},
		},
		Views: []*View{
			{
				Name:        "aws-cis-view",
				Description: "AWS CIS View",
				Queries: []*Query{
					{
						Name:  "test-query-view",
						Query: "SELECT * FROM my.view",
					},
				},
			},
		},
	}

	assert.EqualValues(t, exp, policiesWrapper.Policies[0])
}
