resource "azurerm_resource_group" "iot" {
  name     = "${var.prefix}-iot"
  location = "East US"
}

resource "azurerm_storage_account" "iot" {
  name                     = "${var.prefix}cqiot"
  resource_group_name      = azurerm_resource_group.iot.name
  location                 = azurerm_resource_group.iot.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "iot" {
  name                  = "iot"
  storage_account_name  = azurerm_storage_account.iot.name
  container_access_type = "private"
}

resource "azurerm_eventhub_namespace" "example" {
  name                = "${var.prefix}-iot"
  resource_group_name = azurerm_resource_group.iot.name
  location            = azurerm_resource_group.iot.location
  sku                 = "Basic"
}

resource "azurerm_eventhub" "example" {
  name                = "${var.prefix}-iot"
  resource_group_name = azurerm_resource_group.iot.name
  namespace_name      = azurerm_eventhub_namespace.example.name
  partition_count     = 2
  message_retention   = 1
}

resource "azurerm_eventhub_authorization_rule" "example" {
  resource_group_name = azurerm_resource_group.iot.name
  namespace_name      = azurerm_eventhub_namespace.example.name
  eventhub_name       = azurerm_eventhub.example.name
  name                = "${var.prefix}iot"
  send                = true
}

resource "azurerm_iothub" "example" {
  name                = "${var.prefix}-cq-iothub"
  resource_group_name = azurerm_resource_group.iot.name
  location            = azurerm_resource_group.iot.location

  sku {
    name     = "S1"
    capacity = "1"
  }

  endpoint {
    type                       = "AzureIotHub.StorageContainer"
    connection_string          = azurerm_storage_account.iot.primary_blob_connection_string
    name                       = "export"
    batch_frequency_in_seconds = 60
    max_chunk_size_in_bytes    = 10485760
    container_name             = azurerm_storage_container.iot.name
    encoding                   = "Avro"
    file_name_format           = "{iothub}/{partition}_{YYYY}_{MM}_{DD}_{HH}_{mm}"
  }

  endpoint {
    type              = "AzureIotHub.EventHub"
    connection_string = azurerm_eventhub_authorization_rule.example.primary_connection_string
    name              = "export2"
  }

  route {
    name           = "export"
    source         = "DeviceMessages"
    condition      = "true"
    endpoint_names = ["export"]
    enabled        = true
  }

  route {
    name           = "export2"
    source         = "DeviceMessages"
    condition      = "true"
    endpoint_names = ["export2"]
    enabled        = true
  }

  enrichment {
    key            = "tenant"
    value          = "$twin.tags.Tenant"
    endpoint_names = ["export", "export2"]
  }

  cloud_to_device {
    max_delivery_count = 30
    default_ttl        = "PT1H"
    feedback {
      time_to_live       = "PT1H10M"
      max_delivery_count = 15
      lock_duration      = "PT30S"
    }
  }

  tags = var.tags
}