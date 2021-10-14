package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationResourcesPolicyAssignments(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ResourcesPolicyAssignments(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ResourcesPolicyAssignments().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s%s-policy-assignment", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"sku_tier":         "Free",
					"sku_name":         "A0",
					"enforcement_mode": "Default",
					"type":             "Microsoft.Authorization/policyAssignments",
				},
			}},
		}
	})
}
