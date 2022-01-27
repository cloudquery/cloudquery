resource "azurerm_security_center_contact" "security_center_contact" {
  email = "cq-int-tests@example.com"
  phone = "+1-555-555-5555"

  alert_notifications = true
  alerts_to_admins    = true
}

resource "azurerm_security_center_subscription_pricing" "security_pricing" {
  tier          = "Free"
  resource_type = "VirtualMachines"
}

resource "azurerm_security_center_setting" "example" {
  setting_name = "MCAS"
  enabled      = true
}
