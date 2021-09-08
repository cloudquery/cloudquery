package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationLoadBalancers(t *testing.T) {
	testIntegrationHelper(t, resources.LoadBalancers(), []string{"do_loadbalancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.LoadBalancers().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("do-loadbalancer-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do-loadbalancer-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_load_balancer_forwarding_rules",
					ForeignKeyName: "load_balancer_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"entry_port":      1245.0,
								"entry_protocol":  "http",
								"target_port":     3030.0,
								"target_protocol": "http",
							},
						},
					},
				},
			},
		}
	})
}
