package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationNetworkVirtualNetworks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.NetworkVirtualNetworks(), []string{
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.NetworkVirtualNetworks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-network", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"provisioning_state":     "Succeeded",
					"enable_ddos_protection": false,
					"type":                   "Microsoft.Network/virtualNetworks",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_network_virtual_network_peerings",
					ForeignKeyName: "virtual_network_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"allow_virtual_network_access":          true,
							"allow_forwarded_traffic":               false,
							"allow_gateway_transit":                 false,
							"use_remote_gateways":                   false,
							"remote_address_space_address_prefixes": []interface{}{"10.1.0.0/16"},
							"peering_state":                         "Initiated",
							"provisioning_state":                    "Succeeded",
							"name":                                  fmt.Sprintf("%s-%s-peering", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "azure_network_virtual_network_subnets",
					ForeignKeyName: "virtual_network_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"provisioning_state":                    "Succeeded",
							"private_endpoint_network_policies":     "Disabled",
							"private_link_service_network_policies": "Enabled",
							"name":                                  fmt.Sprintf("%s-%s-internal", res.Prefix, res.Suffix),
						},
					}},
				},
			},
		}
	})
}
