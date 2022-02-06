terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
    }
  }
  backend "azurerm" {
    resource_group_name  = "tfstate"
    storage_account_name = "cqprovidertazuretfstate"
    container_name       = "tfstate"
    key                  = "prod.terraform.tfstate"
  }
}
