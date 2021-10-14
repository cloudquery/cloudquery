resource "azurerm_virtual_network" "network1" {
  name = "${var.test_prefix}-${var.test_suffix}-network"
  address_space = [
    "10.0.0.0/16"
  ]
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
}

resource "azurerm_virtual_network" "network2" {
  name = "${var.test_prefix}-${var.test_suffix}-network2"
  address_space = [
    "10.1.0.0/16"
  ]
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
}

resource "azurerm_subnet" "internal" {
  name                                           = "${var.test_prefix}-${var.test_suffix}-internal"
  resource_group_name                            = azurerm_resource_group.resource_group.name
  virtual_network_name                           = azurerm_virtual_network.network1.name
  enforce_private_link_endpoint_network_policies = true
  address_prefixes                               = ["10.0.2.0/24"]
  service_endpoints                              = ["Microsoft.Storage", "Microsoft.Sql"]
}

resource "azurerm_network_watcher" "network_watcher" {
  name                = "${var.test_prefix}-${var.test_suffix}-nw"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  tags = {
    test = "test"
  }
}

resource "azurerm_virtual_network_peering" "network_peering" {
  name                      = "${var.test_prefix}-${var.test_suffix}-peering"
  resource_group_name       = azurerm_resource_group.resource_group.name
  virtual_network_name      = azurerm_virtual_network.network1.name
  remote_virtual_network_id = azurerm_virtual_network.network2.id
}






