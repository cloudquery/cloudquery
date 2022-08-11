resource "azurerm_resource_group" "redis" {
  name     = "${var.prefix}-redis"
  location = "East US"
}

resource "azurerm_redis_cache" "redis_service" {
  name                = "${var.prefix}-cq-redis"
  location            = azurerm_resource_group.redis.location
  resource_group_name = azurerm_resource_group.redis.name
  capacity            = 1
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = false
  minimum_tls_version = "1.2"
  public_network_access_enabled = false

  redis_configuration {
  }
  tags = var.tags
}
