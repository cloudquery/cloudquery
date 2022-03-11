resource "azurerm_resource_group" "sql" {
  name     = "${var.prefix}-sql"
  location = "East US"
}