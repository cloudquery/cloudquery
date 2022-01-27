resource "azurerm_storage_account" "storage_accounts_storage_account" {
  name                     = "sacqinttests"
  resource_group_name      = azurerm_resource_group.cq_int_tests.name
  location                 = azurerm_resource_group.cq_int_tests.location
  account_tier             = "Standard"
  account_replication_type = "GRS"

  blob_properties {
    cors_rule {
      allowed_headers    = ["*"]
      allowed_methods    = ["GET", "HEAD", "POST", "PUT"]
      allowed_origins    = ["https://example.com"]
      exposed_headers    = ["*"]
      max_age_in_seconds = 3600
    }
  }
}

resource "azurerm_storage_account_network_rules" "storage_accounts_permit_subnet" {
  resource_group_name  = azurerm_resource_group.cq_int_tests.name
  storage_account_name = azurerm_storage_account.storage_accounts_storage_account.name

  default_action             = "Allow"
  ip_rules                   = ["187.67.86.15"]
  virtual_network_subnet_ids = module.test_vnet.vnet_subnets
  bypass                     = ["AzureServices", "Metrics"]
  private_link_access {
    endpoint_resource_id = azurerm_private_endpoint.storage_accounts_private_endpoint.subnet_id
  }
}

resource "azurerm_private_endpoint" "storage_accounts_private_endpoint" {
  name                = "pe-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  subnet_id           = module.test_vnet.vnet_subnets[0]

  private_service_connection {
    name                           = "psc-cq-int-tests"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_storage_account.storage_accounts_storage_account.id
    subresource_names              = ["file"]
  }
}

resource "azurerm_storage_container" "storage_container" {
  name                  = "storage-container-cq-int-tests"
  storage_account_name  = azurerm_storage_account.storage_accounts_storage_account.name
  container_access_type = "private"
}
