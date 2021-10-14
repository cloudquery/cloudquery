resource "azurerm_app_service_plan" "web_apps_service_plan" {
  name                = "sp-${var.test_prefix}${var.test_suffix}"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name

  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_app_service" "web_apps_app_service" {
  name                = "as-${var.test_prefix}${var.test_suffix}"
  location            = azurerm_resource_group.resource_group.location
  resource_group_name = azurerm_resource_group.resource_group.name
  app_service_plan_id = azurerm_app_service_plan.web_apps_service_plan.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"

  }

  app_settings = {
    "SOME_KEY" = "some-value"

    min_tls_version = "1.2"
  }

  connection_string {
    name  = "Databaasdfse"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }

  identity {
    type = "SystemAssigned"
  }

  auth_settings {
    enabled          = true
    default_provider = "MicrosoftAccount"
    microsoft {
      client_id     = "x"
      client_secret = "x"
    }
  }

  tags = {
    test = "test"
  }
}