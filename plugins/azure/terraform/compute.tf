module "linux_server" {
    source = "Azure/compute/azurerm"
    resource_group_name = azurerm_resource_group.cq_int_tests.name
    vnet_subnet_id = module.test_vnet.vnet_subnets[0]
    vm_os_simple = "UbuntuServer"
    tags = {
        name = "value"
    }
    depends_on = [azurerm_resource_group.cq_int_tests]
}

resource "azurerm_managed_disk" "compute_disks_disk" {
    name                 = "mydisk"
    location             = azurerm_resource_group.cq_int_tests.location
    resource_group_name  = azurerm_resource_group.cq_int_tests.name
    storage_account_type = "Standard_LRS"
    create_option        = "Empty"
    disk_size_gb         = "1"
    tags = {
        name = "value"
    }
}
