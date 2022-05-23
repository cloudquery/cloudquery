resource "google_service_account" "example" {
  account_id   = "${lower(var.prefix)}-example-gke-service-account"
  display_name = "${lower(var.prefix)}-example-gke-service-account"
  project      = var.project_id
}