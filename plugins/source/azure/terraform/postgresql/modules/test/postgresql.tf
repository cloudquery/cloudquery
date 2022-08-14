resource "random_password" "postgresql" {
  length           = 16
  special          = true
}

resource "azurerm_resource_group" "postgresql" {
  name     = "${var.prefix}-postgresql"
  location = "East US"
}

resource "azurerm_postgresql_server" "test" {
  name                = "${var.prefix}-postgresql"
  location            = azurerm_resource_group.postgresql.location
  resource_group_name = azurerm_resource_group.postgresql.name

  sku_name = "B_Gen5_2"

  storage_mb                   = 5120
  backup_retention_days        = 7
  geo_redundant_backup_enabled = false
  auto_grow_enabled            = true

  administrator_login          = "psqladmin"
  administrator_login_password = random_password.postgresql.result
  version                      = "9.5"
  ssl_enforcement_enabled      = true
}

resource "azurerm_postgresql_database" "test" {
  name                = "${var.prefix}-postgresql-db"
  resource_group_name = azurerm_resource_group.postgresql.name
  server_name         = azurerm_postgresql_server.test.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
}

resource "azurerm_postgresql_firewall_rule" "example" {
  name                = "${var.prefix}pgfirewallrule"
  resource_group_name = azurerm_resource_group.postgresql.name
  server_name         = azurerm_postgresql_server.test.name
  start_ip_address    = "10.0.0.0"
  end_ip_address      = "10.0.0.0"
}