package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecurityAutoProvisioningSettings(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SecurityAutoProvisioningSettings(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.SecurityAutoProvisioningSettings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"auto_provision": "On"})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"auto_provision": "On",
					"resource_type":  "Microsoft.Security/autoProvisioningSettings",
				},
			}},
		}
	})
}
