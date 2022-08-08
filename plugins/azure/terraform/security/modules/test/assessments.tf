resource "azurerm_resource_group" "security" {
  name     = "${var.prefix}-security"
  location = "East US"
}

resource "random_password" "security" {
  length           = 16
  special          = true
}

resource "azurerm_virtual_network" "example" {
  name                = "${var.prefix}-security"
  resource_group_name = azurerm_resource_group.security.name
  location            = azurerm_resource_group.security.location
  address_space       = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "internal" {
  name                 = "${var.prefix}-security"
  resource_group_name  = azurerm_resource_group.security.name
  virtual_network_name = azurerm_virtual_network.example.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_linux_virtual_machine_scale_set" "example" {
  name                = "${var.prefix}-security"
  resource_group_name = azurerm_resource_group.security.name
  location            = azurerm_resource_group.security.location
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"
  admin_password      = random_password.security.result
  disable_password_authentication = false

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.internal.id
    }
  }
}

resource "azurerm_security_center_assessment_policy" "example" {
  display_name = "${var.prefix}-security"
  severity     = "Medium"
  description  = "Test Description"
  categories = ["Compute"]
  remediation_description = "Some description"
  threats = ["DenialOfService"]
  user_impact = "Low"
}

resource "azurerm_security_center_assessment" "example" {
  assessment_policy_id = azurerm_security_center_assessment_policy.example.id
  target_resource_id   = azurerm_linux_virtual_machine_scale_set.example.id

  status {
    code = "Healthy"
  }
}