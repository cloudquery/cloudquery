resource "azurerm_security_center_subscription_pricing" "security_pricing" {
  tier          = "Free"
  resource_type = "VirtualMachines"
}