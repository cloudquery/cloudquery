package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationVpcs(t *testing.T) {
	testIntegrationHelper(t, resources.Vpcs(), []string{"do_volumes.tf", "do_droplets.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Vpcs().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("dovpc%s", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":        fmt.Sprintf("dovpc%s", res.Suffix),
						"ip_range":    "10.10.10.0/24",
						"region_slug": "nyc3",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_vpc_members",
					ForeignKeyName: "vpc_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name": fmt.Sprintf("do-droplet%s-%s", res.Prefix, res.Suffix),
							},
						},
					},
				},
			},
		}
	})
}
