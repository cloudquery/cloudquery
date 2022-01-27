resource "azurerm_mysql_server" "mysql_server_1" {
  name                = "mysql-server-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name

  administrator_login          = "mysqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "GP_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = false
  public_network_access_enabled     = true
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}

resource "azurerm_mysql_configuration" "mysql-config-1" {
  name                = "interactive_timeout"
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  server_name         = azurerm_mysql_server.mysql_server_1.name
  value               = "600"
}

resource "azurerm_private_endpoint" "pe_mysql_1" {
  name                = "pe-mysql-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  subnet_id           = module.test_vnet.vnet_subnets[0]

  private_service_connection {
    name                           = "psc-cq-int-tests"
    private_connection_resource_id = azurerm_mysql_server.mysql_server_1.id
    subresource_names              = ["mysqlServer"]
    is_manual_connection           = false
  }
}
