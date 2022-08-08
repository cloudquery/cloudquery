resource "azurerm_storage_account" "to_monitor" {
  name                     = "${var.prefix}cqtomonitor"
  resource_group_name      = azurerm_resource_group.monitor.name
  location                 = azurerm_resource_group.monitor.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_monitor_action_group" "main" {
  name                = "${var.prefix}-cq-actiongroup"
  resource_group_name = azurerm_resource_group.monitor.name
  short_name          = "p0action"

  webhook_receiver {
    name        = "callmyapi"
    service_uri = "http://cloudquery.io/alert"
  }
}

resource "azurerm_monitor_activity_log_alert" "main" {
  name                = "${var.prefix}-cq-activity-log-alert"
  resource_group_name = azurerm_resource_group.monitor.name
  scopes              = [azurerm_resource_group.monitor.id]
  description         = "This alert will monitor a specific storage account updates."

  criteria {
    resource_id    = azurerm_storage_account.to_monitor.id
    operation_name = "Microsoft.Storage/storageAccounts/write"
    category       = "Recommendation"
  }

  action {
    action_group_id = azurerm_monitor_action_group.main.id

    webhook_properties = {
      from = "terraform"
    }
  }
}