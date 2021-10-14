package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeVirtualMachines(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ComputeVirtualMachines(), []string{
		"azure_compute_virtual_machines.tf",
		"azure_keyvault_vaults.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeVirtualMachines().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-vm", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"hardware_profile_vm_size":       "Standard_B1ls",
					"computer_name":                  "hostname",
					"admin_username":                 "testadmin",
					"type":                           "Microsoft.Compute/virtualMachines",
					"allow_extension_operations":     true,
					"require_guest_provision_signal": true,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_compute_virtual_machine_network_interfaces",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"network_interface_reference_properties_primary": false,
						},
					}},
				},
				{
					Name:           "azure_compute_virtual_machine_resources",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name": fmt.Sprintf("vm-extension-%s-%s", res.Prefix, res.Suffix),
							"type": "Microsoft.Compute/virtualMachines/extensions",
							"virtual_machine_extension_properties": map[string]interface{}{
								"type": "CustomScript",
								"settings": map[string]interface{}{
									"commandToExecute": "hostname && uptime",
								},
								"publisher":               "Microsoft.Azure.Extensions",
								"typeHandlerVersion":      "2.0",
								"autoUpgradeMinorVersion": false,
							},
							"tags": map[string]interface{}{
								"test": "test",
							},
						},
					}},
				},
				{
					Name:           "azure_compute_virtual_machine_secrets",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data:  map[string]interface{}{},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "azure_compute_virtual_machine_secret_vault_certificates",
							ForeignKeyName: "virtual_machine_secret_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data:  map[string]interface{}{},
							}},
						},
					},
				},
			},
		}
	})
}

func TestIntegrationComputeVirtualMachinesWindows(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ComputeVirtualMachines(), []string{
		"azure_compute_virtual_machines_windows.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeVirtualMachines().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-w-vm", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"hardware_profile_vm_size":       "Standard_B1ls",
					"computer_name":                  res.Suffix,
					"admin_username":                 "adminuser",
					"type":                           "Microsoft.Compute/virtualMachines",
					"allow_extension_operations":     true,
					"require_guest_provision_signal": true,
					"windows_configuration_patch_settings_patch_mode": "AutomaticByOS",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_compute_virtual_machine_win_config_rm_listeners",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"protocol": "Http",
						},
					}},
				},
			},
		}
	})
}
