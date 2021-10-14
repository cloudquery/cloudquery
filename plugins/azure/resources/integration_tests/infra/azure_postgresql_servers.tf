resource "azurerm_postgresql_server" "pgsql_server_1" {
  name                = "pgsqlserver-${var.test_prefix}-${var.test_suffix}"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name

  administrator_login          = "psqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "GP_Gen5_4"
  version    = "11"
  storage_mb = 5120

  backup_retention_days        = 7
  geo_redundant_backup_enabled = true
  auto_grow_enabled            = true

  public_network_access_enabled    = true
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"
}

resource "azurerm_postgresql_configuration" "example" {
  name                = "backslash_quote"
  resource_group_name = azurerm_resource_group.resource_group.name
  server_name         = azurerm_postgresql_server.pgsql_server_1.name
  value               = "on"
}

resource "azurerm_postgresql_firewall_rule" "pgsql_fw_rule_1" {
  name                = "office"
  resource_group_name = azurerm_resource_group.resource_group.name
  server_name         = azurerm_postgresql_server.pgsql_server_1.name
  start_ip_address    = "10.0.99.0"
  end_ip_address      = "10.0.101.255"
}

resource "azurerm_private_endpoint" "pe_pgsql_1" {
  name                = "pe-pgsql-${var.test_prefix}-${var.test_suffix}"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  subnet_id           = azurerm_subnet.internal.id

  private_service_connection {
    name                           = "psc-${var.test_prefix}-${var.test_suffix}"
    private_connection_resource_id = azurerm_postgresql_server.pgsql_server_1.id
    subresource_names              = ["postgresqlServer"]
    is_manual_connection           = false
  }
}
