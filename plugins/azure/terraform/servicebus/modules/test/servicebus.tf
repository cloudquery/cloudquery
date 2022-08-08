resource "azurerm_resource_group" "servicebus" {
  name     = "${var.prefix}-servicebus"
  location = "East US"
}

resource "azurerm_servicebus_namespace" "example" {
  name                = "${var.prefix}-servicebus"
  location            = azurerm_resource_group.servicebus.location
  resource_group_name = azurerm_resource_group.servicebus.name
  sku                 = "Standard"

  tags = var.tags
}


resource "azurerm_servicebus_namespace_authorization_rule" "example" {
  name         = "${var.prefix}-servicebus-ar"
  namespace_id = azurerm_servicebus_namespace.example.id

  listen = true
  send   = true
  manage = false
}


resource "azurerm_servicebus_topic" "azurerm_servicebus_topic" {
  name         = "${var.prefix}-servicebus-topic"
  namespace_id = azurerm_servicebus_namespace.example.id

  enable_partitioning = true
}

resource "azurerm_servicebus_subscription" "azurerm_servicebus_subscription" {
  name               = "${var.prefix}-servicebus-topic-sub"
  topic_id           = azurerm_servicebus_topic.azurerm_servicebus_topic.id
  max_delivery_count = 1
}


resource "azurerm_servicebus_subscription_rule" "azurerm_servicebus_subscription_rule" {
  name            = "${var.prefix}_servicebus_rule"
  subscription_id = azurerm_servicebus_subscription.azurerm_servicebus_subscription.id
  filter_type     = "SqlFilter"
  sql_filter      = "colour = 'red'"
}


resource "azurerm_servicebus_topic_authorization_rule" "azurerm_servicebus_topic_authorization_rule" {
  name     = "${var.prefix}-servicebus-topic-ar"
  topic_id = azurerm_servicebus_topic.azurerm_servicebus_topic.id
  listen   = true
  send     = false
  manage   = false
}