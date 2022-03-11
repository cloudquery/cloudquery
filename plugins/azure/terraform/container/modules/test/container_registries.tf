resource "azurerm_container_registry" "managed_clusters_registry" {
  name                = "${var.prefix}cqacr"
  resource_group_name = azurerm_resource_group.container.name
  location            = azurerm_resource_group.container.location
  sku                 = "Standard"
  admin_enabled       = false
}