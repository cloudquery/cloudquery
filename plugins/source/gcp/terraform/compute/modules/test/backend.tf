#################################################################################
## Compute Backend Helpers - Targets
#################################################################################

resource "google_compute_target_ssl_proxy" "gcp_compute_target_ssl_proxies_proxy" {
  name             = "${var.prefix}-target-ssl-proxy"
  backend_service  = google_compute_backend_service.another-backend-service.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.gcp_compute_managed_ssl_certificates_cert.id
  ]
}

resource "google_compute_target_https_proxy" "gcp_compute_target_https_proxies" {
  name             = "${var.prefix}-target-https-proxy"
  url_map          = google_compute_url_map.url_map.id
  ssl_certificates = []
  ssl_policy = google_compute_ssl_policy.ssl-policy.id
}

resource "google_compute_target_http_proxy" "gcp_compute_target_http_proxies" {
  name             = "${var.prefix}-target-http-proxy"
  url_map          = google_compute_url_map.url_map.id
}
#
#################################################################################
## Compute Backend Module
#################################################################################

resource "google_compute_backend_service" "backend-service" {
  name          = "${var.prefix}-backend-service"
  health_checks = [google_compute_https_health_check.backend-service-health-check.id]
  protocol      = "HTTPS"
}

resource "google_compute_backend_service" "another-backend-service" {
  name     = "${var.prefix}-another-backend-service"
  protocol = "SSL"
}

resource "google_compute_backend_service" "internal-backend-service" {
  name     = "${var.prefix}-internal-backend-service"
  protocol = "HTTP"

  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}
#
resource "google_compute_https_health_check" "backend-service-health-check" {
  name               = "${var.prefix}-health-check"
  request_path       = "/"
  check_interval_sec = 60
  timeout_sec        = 1
}