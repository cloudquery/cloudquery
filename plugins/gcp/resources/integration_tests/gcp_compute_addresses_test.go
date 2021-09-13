package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeAddresses(t *testing.T) {
	table := resources.ComputeAddresses()
	testIntegrationHelper(t, table, []string{"network.tf", "gcp_compute_addresses.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":         fmt.Sprintf("compute-addr-%s%s", res.Prefix, res.Suffix),
						"address":      "10.2.133.133",
						"address_type": "INTERNAL",
						"description":  "my description",
					},
				},
			},
		}
	})
}
