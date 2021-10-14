resource "azurerm_security_center_contact" "security_center_contact" {
  email = "${var.test_prefix}${var.test_suffix}@example.com"
  phone = "+1-555-555-5555"

  alert_notifications = true
  alerts_to_admins    = true
}