resource "google_project_service" "memorystore" {
  project            = var.project_id
  service            = "redis.googleapis.com"
  disable_on_destroy = false
}

module "memorystore" {
  depends_on = [google_project_service.memorystore]
  source     = "terraform-google-modules/memorystore/google"
  version    = "~> 4.4.0"

  name    = "${var.prefix}-memorystore"
  project = var.project_id

  region       = var.region
  tier         = "BASIC"
  auth_enabled = true

  labels = var.labels
}