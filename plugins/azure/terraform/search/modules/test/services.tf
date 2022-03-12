resource "random_password" "mariadb" {
  length           = 16
  special          = true
}

resource "azurerm_resource_group" "search" {
  name     = "${var.prefix}-search"
  location = "East US"
}

resource "azurerm_search_service" "example" {
  name                = "${var.prefix}-search"
  resource_group_name = azurerm_resource_group.search.name
  location            = azurerm_resource_group.search.location
  sku                 = "free"
  tags = var.tags
}