resource "azurerm_resource_group" "data_lake" {
  name     = "${var.prefix}-data-lake"
  location = "East US 2"
}

resource "azurerm_data_lake_store" "test" {
  name                = "${var.prefix}cqdatalakestore"
  resource_group_name = azurerm_resource_group.data_lake.name
  location            = azurerm_resource_group.data_lake.location
  encryption_state    = "Enabled"
  encryption_type     = "ServiceManaged"
  tags = var.tags
}

resource "azurerm_data_lake_store_firewall_rule" "example" {
  name                = "${var.prefix}-datalake-store-ip-range"
  account_name        = azurerm_data_lake_store.test.name
  resource_group_name = azurerm_resource_group.data_lake.name
  start_ip_address    = "10.0.0.0"
  end_ip_address      = "10.0.0.1"
}