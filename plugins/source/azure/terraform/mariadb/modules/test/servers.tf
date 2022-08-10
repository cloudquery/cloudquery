resource "random_password" "mariadb" {
  length           = 16
  special          = true
}

resource "azurerm_resource_group" "mariadb" {
  name     = "${var.prefix}-mariadb"
  location = "East US"
}

resource "azurerm_mariadb_server" "example" {
  name                = "${var.prefix}-maria"
  location            = azurerm_resource_group.mariadb.location
  resource_group_name = azurerm_resource_group.mariadb.name

  administrator_login          = "mariadbadmin"
  administrator_login_password = random_password.mariadb.result

  sku_name   = "GP_Gen5_2"
  storage_mb = 5120
  version    = "10.2"

  auto_grow_enabled             = false
  backup_retention_days         = 7
  geo_redundant_backup_enabled  = false
  public_network_access_enabled = false
  ssl_enforcement_enabled       = true

}

resource "azurerm_mariadb_database" "example" {
  name                = "${var.prefix}cqmariadb"
  resource_group_name = azurerm_resource_group.mariadb.name
  server_name         = azurerm_mariadb_server.example.name
  charset             = "utf8"
  collation           = "utf8_general_ci"
}

resource "azurerm_virtual_network" "mariadb" {
  name                = "${var.prefix}-mariadb-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.mariadb.location
  resource_group_name = azurerm_resource_group.mariadb.name
}

resource "azurerm_subnet" "mariadb" {
  name                 = "${var.prefix}-mariadb-subnet"
  resource_group_name  = azurerm_resource_group.mariadb.name
  virtual_network_name = azurerm_virtual_network.mariadb.name
  address_prefixes     = ["10.0.2.0/24"]

  enforce_private_link_endpoint_network_policies = true
}

resource "azurerm_private_endpoint" "mariadb-private-endpoint" {
  name                = "${var.prefix}-cq-maria-endpoint"
  location            = azurerm_resource_group.mariadb.location
  resource_group_name = azurerm_resource_group.mariadb.name
  subnet_id           = azurerm_subnet.mariadb.id

  private_service_connection {
    name                           = "${var.prefix}-maria-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_mariadb_server.example.id
    subresource_names              = ["mariadbServer"]
  }
}