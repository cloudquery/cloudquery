resource "azurerm_sql_managed_instance" "azurerm_sql_managed_instance" {
  name                         = "cq-int-test-managed-instance"
  resource_group_name          = azurerm_resource_group.cq_int_tests.name
  location                     = azurerm_resource_group.cq_int_tests.location
  administrator_login          = "mradministrator"
  administrator_login_password = "thisIsDog11"
  license_type                 = "BasePrice"
  subnet_id                    = module.test_vnet.vnet_subnets[0]
  sku_name                     = "GP_Gen5"
  vcores                       = 4
  storage_size_in_gb           = 32

  depends_on = [
    module.test_vnet
  ]
}

resource "azurerm_sql_managed_database" "azurerm_sql_managed_database" {
  name                    = "cq-int-test-managed-database"
  sql_managed_instance_id = azurerm_sql_managed_instance.azurerm_sql_managed_instance.id
  location                = azurerm_resource_group.cq_int_tests.location
}