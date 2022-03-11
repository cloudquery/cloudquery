resource "random_password" "sql" {
  length           = 16
  special          = true
}

// Those rules are necessary to provision the cluster successfully

resource "azurerm_network_security_group" "sql-managed-security-group" {
  name                = "${var.prefix}-sql-managed-security-group"
  location            = azurerm_resource_group.sql.location
  resource_group_name = azurerm_resource_group.sql.name
}

resource "azurerm_network_security_rule" "allow_management_inbound" {
  name                        = "${var.prefix}-sql-allow-management-inbound"
  priority                    = 106
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_ranges     = ["9000", "9003", "1438", "1440", "1452"]
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "allow_misubnet_inbound" {
  name                        = "${var.prefix}-sql-allow-misubnet-inbound"
  priority                    = 200
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "10.0.0.0/24"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "allow_health_probe_inbound" {
  name                        = "${var.prefix}-sql-allow-health-probe-inbound"
  priority                    = 300
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "AzureLoadBalancer"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "allow_tds_inbound" {
  name                        = "${var.prefix}-sql-allow-tds-inbound"
  priority                    = 1000
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "1433"
  source_address_prefix       = "VirtualNetwork"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "deny_all_inbound" {
  name                        = "${var.prefix}-sql-deny-all-inbound"
  priority                    = 4096
  direction                   = "Inbound"
  access                      = "Deny"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "allow_management_outbound" {
  name                        = "${var.prefix}-sql-allow-management-outbound"
  priority                    = 102
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_ranges     = ["80", "443", "12000"]
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "allow_misubnet_outbound" {
  name                        = "${var.prefix}-sql-allow-misubnet-outbound"
  priority                    = 200
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "10.0.0.0/24"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_network_security_rule" "deny_all_outbound" {
  name                        = "${var.prefix}-sql-deny-all-outbound"
  priority                    = 4096
  direction                   = "Outbound"
  access                      = "Deny"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.sql.name
  network_security_group_name = azurerm_network_security_group.sql-managed-security-group.name
}

resource "azurerm_virtual_network" "sql-managed" {
  name                = "${var.prefix}-sql-vnet-sql-managed"
  resource_group_name = azurerm_resource_group.sql.name
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.sql.location
}

resource "azurerm_subnet" "sql-managed" {
  name                 = "${var.prefix}-sql-subnet-managed"
  resource_group_name  = azurerm_resource_group.sql.name
  virtual_network_name = azurerm_virtual_network.sql-managed.name
  address_prefixes       = ["10.0.0.0/24"]

  delegation {
    name = "managedinstancedelegation"

    service_delegation {
      name    = "Microsoft.Sql/managedInstances"
      actions = ["Microsoft.Network/virtualNetworks/subnets/join/action", "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action", "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action"]
    }
  }
}

resource "azurerm_subnet_network_security_group_association" "example" {
  subnet_id                 = azurerm_subnet.sql-managed.id
  network_security_group_id = azurerm_network_security_group.sql-managed-security-group.id
}

resource "azurerm_route_table" "sql-managed" {
  name                          = "${var.prefix}-sql-routetable-mi"
  location                      = azurerm_resource_group.sql.location
  resource_group_name           = azurerm_resource_group.sql.name
  disable_bgp_route_propagation = false
  depends_on = [
    azurerm_subnet.sql-managed,
  ]
}

resource "azurerm_subnet_route_table_association" "example" {
  subnet_id      = azurerm_subnet.sql-managed.id
  route_table_id = azurerm_route_table.sql-managed.id
}

resource "azurerm_sql_managed_instance" "example" {
  name                         = "${var.prefix}-sql-managed-instance"
  resource_group_name          = azurerm_resource_group.sql.name
  location                     = azurerm_resource_group.sql.location
  administrator_login          = "mradministrator"
  administrator_login_password = random_password.sql.result
  license_type                 = "BasePrice"
  subnet_id                    = azurerm_subnet.sql-managed.id
  sku_name                     = "GP_Gen5"
  vcores                       = 4
  storage_size_in_gb           = 32

  depends_on = [
    azurerm_subnet_network_security_group_association.example,
    azurerm_subnet_route_table_association.example,
  ]
}

// resource "azurerm_sql_managed_database" "azurerm_sql_managed_database" {
//   name                    = "cq-provider-azure-sql-db"
//   sql_managed_instance_id = azurerm_sql_managed_instance.test.id
//   location                = azurerm_resource_group.test.location
// }