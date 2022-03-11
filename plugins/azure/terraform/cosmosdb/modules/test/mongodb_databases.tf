
resource "azurerm_resource_group" "mongodb" {
  name     = "${var.prefix}-mongodb"
  location = "East US"
}

resource "azurerm_cosmosdb_account" "mongodb" {
  name                = "${var.prefix}-cq-mongodb-account"
  resource_group_name = azurerm_resource_group.mongodb.name
  location = azurerm_resource_group.mongodb.location

  offer_type          = "Standard"
  kind                = "MongoDB"

  enable_automatic_failover = true

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["DELETE", "GET", "HEAD", "MERGE", "POST", "OPTIONS", "PUT", "PATCH"]
    allowed_origins = ["*"]
    exposed_headers = ["*"]
    max_age_in_seconds = 730
  }

  capabilities {
    name = "EnableAggregationPipeline"
  }

  capabilities {
    name = "mongoEnableDocLevelTTL"
  }

  capabilities {
    name = "MongoDBv3.4"
  }

  capabilities {
    name = "EnableMongo"
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 300
    max_staleness_prefix    = 100000
  }

  geo_location {
    location          = azurerm_resource_group.mongodb.location
    failover_priority = 0
  }
  tags = var.tags
}

resource "azurerm_cosmosdb_mongo_database" "test" {
  name                = "${var.prefix}-cq-mongodb"
  resource_group_name = azurerm_cosmosdb_account.mongodb.resource_group_name
  account_name        = azurerm_cosmosdb_account.mongodb.name
  throughput          = 400
  // autoscale_settings {
  //   max_throughput = 4000
  // }
}


resource "azurerm_virtual_network" "mongodb" {
  name                = "${var.prefix}-mongodb-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.mongodb.location
  resource_group_name = azurerm_resource_group.mongodb.name
}

resource "azurerm_subnet" "mongodb" {
  name                 = "${var.prefix}-mongodb-subnet"
  resource_group_name  = azurerm_resource_group.mongodb.name
  virtual_network_name = azurerm_virtual_network.mongodb.name
  address_prefixes     = ["10.0.2.0/24"]

  enforce_private_link_endpoint_network_policies = true
}

resource "azurerm_private_endpoint" "mongodb-private-endpoint" {
  name                = "${var.prefix}-cq-cosmosdb-mongo-private-endpoint"
  location            = azurerm_resource_group.mongodb.location
  resource_group_name = azurerm_resource_group.mongodb.name
  subnet_id           = azurerm_subnet.mongodb.id

  private_service_connection {
    name                           = "tfex-mongodb-connection"
    is_manual_connection           = false
    private_connection_resource_id = azurerm_cosmosdb_account.mongodb.id
    subresource_names              = ["MongoDB"]
  }
}