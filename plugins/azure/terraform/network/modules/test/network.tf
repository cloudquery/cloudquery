resource "azurerm_resource_group" "network" {
  name     = "${var.prefix}-network-rg"
  location = "East US"
}

resource "azurerm_express_route_port" "er_port" {
  name                = "${var.prefix}-er-port"
  resource_group_name = azurerm_resource_group.network.name
  location            = azurerm_resource_group.network.location
  peering_location    = "Equinix-Amsterdam-AM5"
  bandwidth_in_gbps   = 10
  encapsulation       = "Dot1Q"
}

resource "azurerm_express_route_circuit" "er_circuit" {
  name                  = "${var.prefix}-circuit"
  resource_group_name   = azurerm_resource_group.network.name
  location              = azurerm_resource_group.network.location
  service_provider_name = "Equinix"
  peering_location      = "Silicon Valley"
  bandwidth_in_mbps     = 50

  sku {
    tier   = "Standard"
    family = "MeteredData"
  }

  tags = {
    environment = "Test"
  }
}

resource "azurerm_route_filter" "route_filter" {
  name                = "${var.prefix}-route-filter"
  resource_group_name   = azurerm_resource_group.network.name
  location              = azurerm_resource_group.network.location

  rule {
    name        = "rule"
    access      = "Allow"
    rule_type   = "Community"
    communities = ["12076:52004"]
  }
}

# express route gateway
resource "azurerm_virtual_wan" "virtual_wan" {
  name                = "${var.prefix}-er-virtualwan"
  resource_group_name = azurerm_resource_group.network.name
  location            = azurerm_resource_group.network.location
}

resource "azurerm_virtual_hub" "virtual_hub" {
  name                = "${var.prefix}-er-virtualhub"
  resource_group_name = azurerm_resource_group.network.name
  location            = azurerm_resource_group.network.location
  virtual_wan_id      = azurerm_virtual_wan.virtual_wan.id
  address_prefix      = "10.0.1.0/24"
}

resource "azurerm_express_route_gateway" "gateway" {
  name                = "${var.prefix}-er-gateway"
  resource_group_name = azurerm_resource_group.network.name
  location            = azurerm_resource_group.network.location
  virtual_hub_id      = azurerm_virtual_hub.virtual_hub.id
  scale_units         = 1

  tags = {
    environment = "test"
  }
}
