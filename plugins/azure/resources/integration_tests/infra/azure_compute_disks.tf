resource "azurerm_managed_disk" "compute_disks_disk" {
  name                 = "disks-${var.test_prefix}-${var.test_suffix}"
  location             = azurerm_resource_group.resource_group.location
  resource_group_name  = azurerm_resource_group.resource_group.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"
  tags = {
    environment = "staging"
  }
}



