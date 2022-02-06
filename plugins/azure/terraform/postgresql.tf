module "postgresql" {
  source = "Azure/postgresql/azurerm"

  resource_group_name = azurerm_resource_group.cq_int_tests.name
  location            = azurerm_resource_group.cq_int_tests.location

  server_name                  = "cq-provider-azure-pgsql"
  sku_name                     = "GP_Gen5_2"
  storage_mb                   = 5120
  backup_retention_days        = 7
  geo_redundant_backup_enabled = false
  administrator_login          = "psqladminun"
  administrator_password       = random_password.password.result
  server_version               = "11"
  ssl_enforcement_enabled      = true
  db_names                     = ["my_db1", "my_db2"]
  db_charset                   = "UTF8"
  db_collation                 = "English_United States.1252"

  firewall_rule_prefix = "firewall-"
  firewall_rules = [
    { name = "test1", start_ip = "10.0.0.5", end_ip = "10.0.0.8" },
    { start_ip = "127.0.0.0", end_ip = "127.0.1.0" },
  ]

  tags = {
    Environment = "Production",
  }

  postgresql_configurations = {
    backslash_quote = "on",
  }
}
