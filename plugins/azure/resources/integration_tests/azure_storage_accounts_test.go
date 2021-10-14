package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationStorageAccounts(t *testing.T) {
	awsTestIntegrationHelper(t, resources.StorageAccounts(), []string{
		"azure_storage_accounts.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.StorageAccounts().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%stsa", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"sku_name":            "Standard_GRS",
					"sku_tier":            "Standard",
					"kind":                "StorageV2",
					"minimum_tls_version": "TLS1_0",
					"type":                "Microsoft.Storage/storageAccounts",
					"identity_type":       "None",
					"status_of_primary":   "available",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_storage_account_network_rule_set_ip_rules",
					ForeignKeyName: "account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"ip_address_or_range": "187.67.86.15",
							"action":              "Allow",
						},
					}},
				},
				{
					Name:           "azure_storage_account_network_rule_set_virtual_network_rules",
					ForeignKeyName: "account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"action": "Allow",
							"state":  "Succeeded",
						},
					}},
				},
				{
					Name:           "azure_storage_account_private_endpoint_connections",
					ForeignKeyName: "account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"private_link_service_connection_state_status":          "Approved",
							"private_link_service_connection_state_description":     "Auto-Approved",
							"private_link_service_connection_state_action_required": "None",
							"provisioning_state": "Succeeded",
							"type":               "Microsoft.Storage/storageAccounts/privateEndpointConnections",
						},
					}},
				},
				{
					Name:           "azure_storage_blob_services",
					ForeignKeyName: "account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"delete_retention_policy_enabled": false,
							"sku_name":                        "Standard_GRS",
							"sku_tier":                        "Standard",
							"name":                            "default",
							"type":                            "Microsoft.Storage/storageAccounts/blobServices",
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{{
						Name:           "azure_storage_blob_service_cors_rules",
						ForeignKeyName: "blob_service_cq_id",
						ExpectedValues: []providertest.ExpectedValue{{
							Count: 1,
							Data: map[string]interface{}{
								"allowed_origins":    []interface{}{"https://example.com"},
								"allowed_methods":    []interface{}{"GET", "HEAD", "POST", "PUT"},
								"max_age_in_seconds": float64(3600),
								"exposed_headers":    []interface{}{"*"},
								"allowed_headers":    []interface{}{"*"},
							},
						}},
					}},
				},
				{
					Name:           "azure_storage_containers",
					ForeignKeyName: "account_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"public_access": "None",
							"name":          fmt.Sprintf("storage-container-%s%s", res.Prefix, res.Suffix),
							"type":          "Microsoft.Storage/storageAccounts/blobServices/containers",
						},
					}},
				},
			},
		}
	})
}
