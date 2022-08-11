resource "azurerm_resource_group" "keyvault" {
  name     = "${var.prefix}-keyvault"
  location = "East US"
}

resource "azurerm_key_vault" "azurerm_key_vault" {
  name                        = "${var.prefix}vault"
  location                    = azurerm_resource_group.keyvault.location
  resource_group_name         = azurerm_resource_group.keyvault.name
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  sku_name                    = "standard"
  soft_delete_retention_days  = 8
  enabled_for_disk_encryption = true
  purge_protection_enabled    = false
  enabled_for_template_deployment = true
  enabled_for_deployment      = true
  enable_rbac_authorization = true
  tags = var.tags
}


resource "azurerm_key_vault_key" "generated" {
  name         = "${var.prefix}key"
  key_vault_id = azurerm_key_vault.azurerm_key_vault.id
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
  tags = var.tags
}
