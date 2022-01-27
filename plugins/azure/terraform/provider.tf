provider "azurerm" {
    features {}
}

data "azurerm_subscription" "current" {}
data "azurerm_client_config" "current" {}
