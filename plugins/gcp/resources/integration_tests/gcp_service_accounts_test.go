package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamServiceAccounts(t *testing.T) {
	testIntegrationHelper(t, resources.IamServiceAccounts(), []string{
		"service-account.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.IamServiceAccounts().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"display_name": fmt.Sprintf("Service Account  %s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"disabled": false,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_iam_service_account_keys",
					ForeignKeyName: "service_account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Data: map[string]interface{}{
								"key_origin":    "GOOGLE_PROVIDED",
								"key_algorithm": "KEY_ALG_RSA_2048",
							},
						},
					},
				},
			},
		}
	})
}
