resource "random_password" "compute" {
  length           = 16
  special          = true
}

resource "azurerm_resource_group" "compute" {
  name     = "${var.prefix}-compute"
  location = "East US"
}

// windows machine

resource "azurerm_virtual_network" "compute" {
  name                = "${var.prefix}-compute"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.compute.location
  resource_group_name = azurerm_resource_group.compute.name
}

resource "azurerm_subnet" "compute" {
  name                 = "${var.prefix}-compute-internal"
  resource_group_name  = azurerm_resource_group.compute.name
  virtual_network_name = azurerm_virtual_network.compute.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_network_interface" "windows-nic" {
  name                = "${var.prefix}-compute-win"
  location            = azurerm_resource_group.compute.location
  resource_group_name = azurerm_resource_group.compute.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.compute.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_windows_virtual_machine" "example" {
  name                = "${var.prefix}-compute-win"
  resource_group_name = azurerm_resource_group.compute.name
  location            = azurerm_resource_group.compute.location
  size                = "Standard_F2"
  admin_username      = "adminuser"
  admin_password      = random_password.compute.result
  network_interface_ids = [
    azurerm_network_interface.windows-nic.id,
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2016-Datacenter"
    version   = "latest"
  }
}

// resource "azurerm_virtual_machine_extension" "linextension" {
//   name                 = "linextension"
//   virtual_machine_id   = azurerm_virtual_machine.main.id
//   publisher            = "Microsoft.Azure.Extensions"
//   type                 = "CustomScript"
//   type_handler_version = "2.0"

//   settings = <<SETTINGS
//     {
//         "commandToExecute": "hostname"
//     }
// SETTINGS
//   protected_settings = <<SETTINGS
//     {}
// SETTINGS
//   tags = var.tags
// }
