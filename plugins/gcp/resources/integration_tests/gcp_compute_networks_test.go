package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeNetworks(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeNetworks(), []string{"network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeNetworks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("network-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                        fmt.Sprintf("network-%s%s", res.Prefix, res.Suffix),
						"routing_config_routing_mode": "REGIONAL",
						"auto_create_subnetworks":     true,
						"kind":                        "compute#network",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_network_peerings",
					ForeignKeyName: "network_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":                                fmt.Sprintf("network-peering-%s%s", res.Prefix, res.Suffix),
								"auto_create_routes":                  true,
								"exchange_subnet_routes":              true,
								"export_custom_routes":                false,
								"export_subnet_routes_with_public_ip": true,
								"import_custom_routes":                false,
								"import_subnet_routes_with_public_ip": false,
								"state":                               "INACTIVE",
								"peer_mtu":                            float64(0),
								"network_name":                        fmt.Sprintf("network-%s%s", res.Prefix, res.Suffix),
							},
						},
					},
				},
			},
		}
	})
}
