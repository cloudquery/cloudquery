resource "azurerm_resource_group" "monitor" {
  name     = "${var.prefix}-monitor"
  location = "East US"
}