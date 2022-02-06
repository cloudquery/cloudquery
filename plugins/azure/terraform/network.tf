module "test_vnet" {
    source = "Azure/vnet/azurerm"
    version = "= 2.5.0"
    resource_group_name = azurerm_resource_group.cq_int_tests.name
    subnet_prefixes = ["10.0.1.0/24"]
    subnet_names = ["subnet1"]
    subnet_enforce_private_link_endpoint_network_policies = {
      "subnet1" = true
    }
    subnet_service_endpoints = {
      "subnet1" = ["Microsoft.Storage", "Microsoft.Sql"]
    }
    depends_on = [azurerm_resource_group.cq_int_tests]
}

resource "azurerm_public_ip" "public_ips_ip" {
  name                = "public-ip-cq-int-tests"
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  location            = azurerm_resource_group.cq_int_tests.location
  allocation_method   = "Static"

  tags = {
    environment = "Production"
  }
}

