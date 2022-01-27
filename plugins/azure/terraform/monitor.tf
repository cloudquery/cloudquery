resource "azurerm_monitor_action_group" "log_alerts_action_group" {
  name                = "mag-cq-int-tests"
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  short_name          = "p0action"

  webhook_receiver {
    name        = "callmyapi"
    service_uri = "http://example.com/alert"
  }
}

resource "azurerm_monitor_activity_log_alert" "log_alerts_log_alert" {
  name                = "mala-cq-int-tests"
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  scopes              = [azurerm_resource_group.cq_int_tests.id]
  description         = "This alert will monitor a specific storage account updates."

  criteria {
    resource_id    = azurerm_storage_account.storage_accounts_storage_account.id
    operation_name = "Microsoft.Storage/storageAccounts/write"
    category       = "Recommendation"
  }

  tags = {
    test = "test"
  }

  action {
    action_group_id = azurerm_monitor_action_group.log_alerts_action_group.id

    webhook_properties = {
      from = "terraform"
    }
  }
}

resource "azurerm_monitor_diagnostic_setting" "diagnostic_settings_ds" {
  name               = "mds-cq-int-tests"
  target_resource_id = module.test_vnet.vnet_id
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

resource "azurerm_eventhub_namespace" "log_profiles_eventhub" {
  name                = "en-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  sku                 = "Standard"
  capacity            = 2
}

resource "azurerm_monitor_log_profile" "log_profiles_log_profile" {
  name = "mlp-cq-int-tests"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  # RootManageSharedAccessKey is created by default with listen, send, manage permissions
  servicebus_rule_id = "${azurerm_eventhub_namespace.log_profiles_eventhub.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.storage_accounts_storage_account.id

  retention_policy {
    enabled = true
    days    = 7
  }
}
