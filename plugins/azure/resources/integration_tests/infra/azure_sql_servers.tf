resource "azurerm_mssql_server_extended_auditing_policy" "mssql_server_au_policy" {
  server_id                               = azurerm_mssql_server.mssql_server_1.id
  storage_endpoint                        = azurerm_storage_account.storage_accounts_storage_account.primary_blob_endpoint
  storage_account_access_key              = azurerm_storage_account.storage_accounts_storage_account.primary_access_key
  storage_account_access_key_is_secondary = false
  retention_in_days                       = 6
}

resource "azurerm_mssql_server" "mssql_server_1" {
  name                         = "mssql-1-${var.test_prefix}-${var.test_suffix}"
  location                     = azurerm_resource_group.resource_group.location
  resource_group_name          = azurerm_resource_group.resource_group.name
  version                      = "12.0"
  administrator_login          = "missadministrator"
  administrator_login_password = "thisIsKat11"
  minimum_tls_version          = "1.2"
}

resource "azurerm_mssql_firewall_rule" "mssql_fw_rule_1" {
  name             = "office"
  server_id        = azurerm_mssql_server.mssql_server_1.id
  start_ip_address = "10.0.99.0"
  end_ip_address   = "10.0.101.255"
}

resource "azurerm_private_endpoint" "pe_mssql_1" {
  name                = "pe-mssql-${var.test_prefix}-${var.test_suffix}"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  subnet_id           = azurerm_subnet.internal.id

  private_service_connection {
    name                           = "psc-${var.test_prefix}-${var.test_suffix}"
    private_connection_resource_id = azurerm_mssql_server.mssql_server_1.id
    subresource_names              = ["sqlServer"]
    is_manual_connection           = false
  }
}

resource "azurerm_mssql_database_extended_auditing_policy" "mssql_db_au_policy" {
  database_id                             = azurerm_mssql_database.mssql_db_1.id
  storage_endpoint                        = azurerm_storage_account.storage_accounts_storage_account.primary_blob_endpoint
  storage_account_access_key              = azurerm_storage_account.storage_accounts_storage_account.primary_access_key
  storage_account_access_key_is_secondary = false
  retention_in_days                       = 6
}

resource "azurerm_mssql_database" "mssql_db_1" {
  name           = "mssql-db-1-${var.test_prefix}-${var.test_suffix}"
  server_id      = azurerm_mssql_server.mssql_server_1.id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  max_size_gb    = 4
  read_scale     = true
  sku_name       = "BC_Gen5_2"
  zone_redundant = true
}
