package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationNetworkSecurityGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.NetworkSecurityGroups(), []string{
		"azure_storage_accounts.tf",
		"azure_network_security_groups.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.NetworkSecurityGroups().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-nsg", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"provisioning_state": "Succeeded",
					"type":               "Microsoft.Network/networkSecurityGroups",
					"tags":               map[string]interface{}{"environment": "Production"},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_network_security_group_default_security_rules",
					ForeignKeyName: "security_group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"protocol":                   "*",
							"source_port_range":          "*",
							"destination_port_range":     "*",
							"source_address_prefix":      "AzureLoadBalancer",
							"destination_address_prefix": "*",
							"access":                     "Allow",
							"priority":                   float64(65001),
							"direction":                  "Inbound",
							"provisioning_state":         "Succeeded",
							"name":                       "AllowAzureLoadBalancerInBound",
							"type":                       "Microsoft.Network/networkSecurityGroups/defaultSecurityRules",
						},
					}},
				},
				{
					Name:           "azure_network_security_group_flow_logs",
					ForeignKeyName: "security_group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"enabled":                              true,
							"retention_policy_days":                float64(7),
							"retention_policy_enabled":             true,
							"format_type":                          "JSON",
							"format_version":                       float64(1),
							"flow_analytics_configuration_enabled": true,
							"flow_analytics_configuration_traffic_analytics_interval": float64(10),
						},
					}},
				},
				{
					Name:           "azure_network_security_group_security_rules",
					ForeignKeyName: "security_group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"protocol":                   "Tcp",
							"source_port_range":          "3389",
							"destination_port_range":     "*",
							"source_address_prefix":      "*",
							"destination_address_prefix": "*",
							"access":                     "Allow",
							"priority":                   float64(121),
							"direction":                  "Inbound",
							"provisioning_state":         "Succeeded",
							"name":                       "test12223",
						},
					}},
				},
			},
		}
	})
}
