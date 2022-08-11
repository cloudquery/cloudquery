resource "azurerm_resource_group" "logic" {
  name     = "${var.prefix}-logic"
  location = "East US"
}

resource "azurerm_logic_app_workflow" "logic" {
  name                = "${var.prefix}-logic"
  location            = azurerm_resource_group.logic.location
  resource_group_name = azurerm_resource_group.logic.name
}
