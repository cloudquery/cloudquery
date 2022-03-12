resource "azurerm_resource_group" "batch" {
  name     = "${var.prefix}-batch"
  location = "East US"
}

resource "azurerm_storage_account" "example" {
  name                     = "cq${var.prefix}batch"
  resource_group_name      = azurerm_resource_group.batch.name
  location                 = azurerm_resource_group.batch.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_batch_account" "example" {
  name                 = "${var.prefix}batch"
  resource_group_name  = azurerm_resource_group.batch.name
  location             = azurerm_resource_group.batch.location
  pool_allocation_mode = "BatchService"
  storage_account_id   = azurerm_storage_account.example.id

  tags = var.tags
}
