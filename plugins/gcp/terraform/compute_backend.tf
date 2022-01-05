################################################################################
# Compute Backend Helpers - Targets
################################################################################

resource "google_compute_target_ssl_proxy" "gcp_compute_target_ssl_proxies_proxy" {
  name             = "${local.prefix}-target-ssl-proxy"
  backend_service  = google_compute_backend_service.another-backend-service.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.gcp_compute_managed_ssl_certificates_cert.id
  ]
}

resource "google_compute_target_https_proxy" "gcp_compute_target_https_proxies" {
  name             = "${local.prefix}-target-https-proxy"
  url_map          = google_compute_url_map.external_url_map.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.gcp_compute_managed_ssl_certificates_cert.id
  ]
}

################################################################################
# Compute Backend Module
################################################################################

resource "google_compute_backend_service" "backend-service" {
  name          = "${local.prefix}-backend-service"
  health_checks = [google_compute_https_health_check.backend-service-health-check.id]
  protocol      = "HTTPS"

  backend {
    group = module.compute_instance_group.instance_group
  }

}

resource "google_compute_backend_service" "another-backend-service" {
  name     = "${local.prefix}-another-backend-service"
  protocol = "SSL"
}

resource "google_compute_backend_service" "internal-backend-service" {
  name     = "${local.prefix}-internal-backend-service"
  protocol = "HTTP"

  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}

resource "google_compute_https_health_check" "backend-service-health-check" {
  name               = "health-check"
  request_path       = "/"
  check_interval_sec = 60
  timeout_sec        = 1
}