resource "azurerm_data_lake_analytics_account" "example" {
  name                = "${var.prefix}cqdatalakeacc"
  resource_group_name = azurerm_resource_group.data_lake.name
  location            = azurerm_resource_group.data_lake.location

  default_store_account_name = azurerm_data_lake_store.test.name
  tags = var.tags
}

resource "azurerm_data_lake_analytics_firewall_rule" "example" {
  name                = "${var.prefix}-datalake-anal-ip-range"
  account_name        = azurerm_data_lake_analytics_account.example.name
  resource_group_name = azurerm_resource_group.data_lake.name
  start_ip_address    = "10.0.0.0"
  end_ip_address      = "10.0.0.1"
}

