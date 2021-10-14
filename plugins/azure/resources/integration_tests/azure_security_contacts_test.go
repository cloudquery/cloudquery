package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecurityContacts(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SecurityContacts(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.SecurityContacts().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"email": fmt.Sprintf("%s%s@example.com", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"phone":               "+1-555-555-5555",
					"alert_notifications": "On",
					"alerts_to_admins":    "On",
					"resource_type":       "Microsoft.Security/securityContacts",
				},
			}},
		}
	})
}
