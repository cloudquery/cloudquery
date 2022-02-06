resource "azurerm_redis_cache" "redis_service" {
  name                = "redis-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name
  capacity            = 1
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = false
  minimum_tls_version = "1.2"
  public_network_access_enabled = false

  redis_configuration {
  }
}
