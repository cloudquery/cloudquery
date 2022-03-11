resource "azurerm_storage_account" "monitor_diagnostic" {
  name                     = "${var.prefix}cqmonitordiag"
  resource_group_name      = azurerm_resource_group.monitor.name
  location                 = azurerm_resource_group.monitor.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  blob_properties {
    versioning_enabled = true
    change_feed_enabled = true
  }

}

resource "azurerm_log_analytics_workspace" "example" {
  name                = "${var.prefix}-cq-log-workspace"
  location            = azurerm_resource_group.monitor.location
  resource_group_name = azurerm_resource_group.monitor.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_key_vault" "to_monitor" {
  name                        = "${var.prefix}cqvaulttomonitor"
  location                    = azurerm_resource_group.monitor.location
  resource_group_name         = azurerm_resource_group.monitor.name
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

resource "azurerm_monitor_diagnostic_setting" "example" {
  name               = "${var.prefix}-cq-diagnostic-settings"
  target_resource_id = azurerm_key_vault.to_monitor.id
  storage_account_id = azurerm_storage_account.monitor_diagnostic.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.example.id
  log {
    category = "AuditEvent"
    enabled  = false

    retention_policy {
      enabled = false
    }
  }

  metric {
    category = "AllMetrics"

    retention_policy {
      enabled = false
    }
  }
}