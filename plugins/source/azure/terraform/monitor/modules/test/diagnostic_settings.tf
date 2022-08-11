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

resource "azurerm_virtual_network" "to_monitor" {
  resource_group_name = azurerm_resource_group.monitor.name
  location            = azurerm_resource_group.monitor.location
  name                = "vnet-test-001"
  address_space       = ["192.168.0.0/16"]
}

data "azurerm_monitor_diagnostic_categories" "to_monitor" {
  resource_id = azurerm_virtual_network.to_monitor.id
}

resource "azurerm_monitor_diagnostic_setting" "example" {
  name               = "${var.prefix}-cq-diagnostic-settings"
  target_resource_id = azurerm_virtual_network.to_monitor.id
  storage_account_id = azurerm_storage_account.monitor_diagnostic.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.example.id
  dynamic "log" {
    for_each = data.azurerm_monitor_diagnostic_categories.to_monitor.logs
    content {
      category = log.value
      retention_policy {
        days    = 0
        enabled = false
      }
    }
  }

  dynamic "metric" {
    for_each = data.azurerm_monitor_diagnostic_categories.to_monitor.metrics
    content {
      category = metric.value
      retention_policy {
        days    = 0
        enabled = false
      }
    }
  }
}