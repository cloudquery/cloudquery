package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationKeyvaultVaults(t *testing.T) {
	awsTestIntegrationHelper(t, resources.KeyvaultVaults(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.KeyvaultVaults().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf(
					"vault-%s%s",
					lastN(res.Prefix, 9),
					lastN(res.Suffix, 9),
				)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"sku_name":                        "standard",
					"sku_family":                      "A",
					"type":                            "Microsoft.KeyVault/vaults",
					"enabled_for_deployment":          true,
					"enabled_for_disk_encryption":     true,
					"enabled_for_template_deployment": false,
					"enable_soft_delete":              true,
					"soft_delete_retention_in_days":   float64(7),
					"enable_rbac_authorization":       false,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_keyvault_vault_access_policies",
					ForeignKeyName: "vault_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"permissions_keys":         []interface{}{"backup", "recover", "create", "decrypt", "delete", "encrypt", "get", "import", "list", "purge", "recover", "restore", "sign", "unwrapKey", "update", "verify", "wrapKey"},
							"permissions_secrets":      []interface{}{"backup", "delete", "recover", "get", "list", "purge", "recover", "restore", "set"},
							"permissions_certificates": []interface{}{"create", "recover", "delete", "deleteissuers", "get", "getissuers", "import", "list", "listissuers", "managecontacts", "manageissuers", "purge", "setissuers", "update"},
							"permissions_storage":      []interface{}{},
						},
					}},
				},
				{
					Name:           "azure_keyvault_vault_keys",
					ForeignKeyName: "vault_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"recoverable_days": float64(7),
							"recovery_level":   "CustomizedRecoverable+Purgeable",
							"enabled":          true,
							"managed":          nil,
							"tags": map[string]interface{}{
								"test": "test",
							},
						},
					}},
				},
				{
					Name:           "azure_keyvault_vault_secrets",
					ForeignKeyName: "vault_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"recoverable_days": float64(7),
							"recovery_level":   "CustomizedRecoverable+Purgeable",
							"enabled":          true,
							"managed":          true,
							"content_type":     "application/x-pkcs12",
						},
					}},
				},
			},
		}
	})
}
