resource "azurerm_resource_group" "search" {
  name     = "${var.prefix}-search"
  location = "East US"
}

resource "azurerm_search_service" "example" {
  name                = "${var.prefix}-search"
  resource_group_name = azurerm_resource_group.search.name
  location            = azurerm_resource_group.search.location
  sku                 = "basic"
  public_network_access_enabled = false
  tags = var.tags
}

resource "azurerm_virtual_network" "search" {
  name                = "${var.prefix}-search"
  resource_group_name = azurerm_resource_group.search.name
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.search.location
}

resource "azurerm_subnet" "search" {
  name                 = "${var.prefix}-search"
  resource_group_name  = azurerm_resource_group.search.name
  virtual_network_name = azurerm_virtual_network.search.name
  address_prefixes       = ["10.0.0.0/24"]
  enforce_private_link_endpoint_network_policies = true
}

resource "azurerm_private_endpoint" "search-private-endpoint" {
  name                = "${var.prefix}-search-private-endpoint"
  location            = azurerm_resource_group.search.location
  resource_group_name = azurerm_resource_group.search.name
  subnet_id           = azurerm_subnet.search.id

  private_service_connection {
    name                           = "${var.prefix}-search-private-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_search_service.example.id
    subresource_names              = ["searchService"]
  }
}