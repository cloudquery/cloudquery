resource "azurerm_monitor_diagnostic_setting" "diagnostic_settings_ds" {
  name               = "${var.test_prefix}-${var.test_suffix}-ds"
  target_resource_id = azurerm_virtual_network.network1.id
  storage_account_id = azurerm_storage_account.storage_accounts_storage_account.id

  log {
    category = "VMProtectionAlerts"
    enabled  = true
    retention_policy {
      days    = 1
      enabled = false
    }
  }

  metric {
    category = "AllMetrics"
    enabled  = true

    retention_policy {
      enabled = false
    }
  }
}