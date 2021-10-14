package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationContainerManagedClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ContainerManagedClusters(), []string{
		"azure_storage_accounts.tf",
		"networks.tf",
		"azure_container_managed_clusters.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ContainerManagedClusters().Name,

			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf(
					"%s-%s-aks",
					res.Prefix,
					res.Suffix,
				)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"provisioning_state":                  "Succeeded",
					"power_state_code":                    "Running",
					"max_agent_pools":                     float64(100),
					"kubernetes_version":                  "1.20.9",
					"dns_prefix":                          fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
					"windows_profile_admin_username":      "azureuser",
					"windows_profile_enable_csi_proxy":    true,
					"service_principal_profile_client_id": "msi",
					"addon_profiles": map[string]interface{}{
						"omsagent":                  map[string]interface{}{"enabled": false},
						"azurepolicy":               map[string]interface{}{"enabled": false},
						"kubeDashboard":             map[string]interface{}{"enabled": false},
						"aciConnectorLinux":         map[string]interface{}{"enabled": false},
						"httpApplicationRouting":    map[string]interface{}{"enabled": false},
						"ingressApplicationGateway": map[string]interface{}{"enabled": false}},
					"enable_rbac":                                              true,
					"network_profile_network_plugin":                           "azure",
					"network_profile_network_policy":                           "azure",
					"network_profile_service_cidr":                             "172.17.0.0/16",
					"network_profile_dns_service_ip":                           "172.17.0.12",
					"network_profile_docker_bridge_cidr":                       "172.18.0.12/16",
					"network_profile_outbound_type":                            "loadBalancer",
					"network_profile_load_balancer_sku":                        "Standard",
					"network_profile_load_balancer_managed_outbound_ips_count": float64(1),
					"disable_local_accounts":                                   false,
					"sku_name":                                                 "Basic",
					"tags": map[string]interface{}{
						"test": "test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_container_managed_cluster_agent_pool_profiles",
					ForeignKeyName: "managed_cluster_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"enable_fips":               false,
							"enable_encryption_at_host": false,
							"node_labels": map[string]interface{}{
								"node-type": "system",
							},
							"tags": map[string]interface{}{
								"test": "test",
							},
							"enable_node_public_ip": false,
							"power_state_code":      "Running",
							"provisioning_state":    "Succeeded",
							"type":                  "VirtualMachineScaleSets",
							"mode":                  "System",
							"min_count":             float64(2),
							"max_count":             float64(2),
							"os_sku":                "Ubuntu",
							"os_type":               "Linux",
							"kubelet_disk_type":     "OS",
							"os_disk_type":          "Managed",
							"os_disk_size_gb":       float64(128),
							"vm_size":               "Standard_B2s",
							"count":                 float64(2),
							"name":                  "default",
						},
					}},
				},
			},
		}
	})
}
