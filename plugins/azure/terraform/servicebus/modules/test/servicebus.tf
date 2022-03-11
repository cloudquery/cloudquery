resource "azurerm_resource_group" "servicebus" {
  name     = "${var.prefix}-servicebus"
  location = "East US"
}

resource "azurerm_servicebus_namespace" "example" {
  name                = "${var.prefix}-servicebus"
  location            = azurerm_resource_group.servicebus.location
  resource_group_name = azurerm_resource_group.servicebus.name
  sku                 = "Standard"

  tags = var.tags
}