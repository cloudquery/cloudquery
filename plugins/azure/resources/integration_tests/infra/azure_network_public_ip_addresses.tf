resource "azurerm_public_ip" "public_ips_ip" {
  name                = "${var.test_prefix}-${var.test_suffix}-ip"
  resource_group_name = azurerm_resource_group.resource_group.name
  location            = azurerm_resource_group.resource_group.location
  allocation_method   = "Static"

  tags = {
    environment = "Production"
  }
}