resource "azurerm_resource_group" "keyvault" {
  name     = "${var.prefix}-keyvault"
  location = "East US"
}

resource "azurerm_key_vault" "test" {
  name                        = "${var.prefix}cqvault"
  location                    = azurerm_resource_group.keyvault.location
  resource_group_name         = azurerm_resource_group.keyvault.name
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
      "List",
      "Get",
    ]

    key_permissions = [
      "purge",
      "List",
      "Create",
      "Get",
    ]

    secret_permissions = [
      "delete",
      "List",
      "Set",
      "Get",
    ]
  }
}


resource "azurerm_key_vault_key" "generated" {
  name         = "${var.prefix}cqkey"
  key_vault_id = azurerm_key_vault.test.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}