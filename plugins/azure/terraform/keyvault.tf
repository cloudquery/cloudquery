resource "azurerm_key_vault" "keyvaults_keyvault" {
  name                        = "kevaultcqproviderazure"
  location                    = azurerm_resource_group.cq_int_tests.location
  resource_group_name         = azurerm_resource_group.cq_int_tests.name
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  sku_name                    = "standard"
  soft_delete_retention_days  = 7
  enabled_for_disk_encryption = true
  purge_protection_enabled    = false
  enabled_for_deployment      = true

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    certificate_permissions = [
      "create",
    ]

    key_permissions = [
      "purge",
    ]

    secret_permissions = [
      "delete",
    ]
  }
}
