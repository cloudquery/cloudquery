resource "google_service_account" "service_account" {
  account_id   = "account-${var.test_suffix}"
  display_name = "Service Account  ${var.test_prefix}${var.test_suffix}"
}
