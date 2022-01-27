resource "azurerm_app_service_plan" "web_apps_service_plan" {
  name                = "sp-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name

  sku {
    tier = "Standard"
    size = "S1"
  }
}

resource "azurerm_app_service" "web_apps_app_service" {
  name                = "as-cq-int-tests"
  location            = azurerm_resource_group.cq_int_tests.location
  resource_group_name = azurerm_resource_group.cq_int_tests.name
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
