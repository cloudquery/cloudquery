package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCoreServiceAccounts(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.CoreServiceAccounts(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_core_service_accounts",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("service-account%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("service-account%s%s", res.Prefix, res.Suffix),
				},
			}},
		}
	})
}
