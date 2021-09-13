resource "google_compute_target_ssl_proxy" "gcp_compute_target_ssl_proxies_proxy" {
  name            = "ssl-proxy-${var.test_prefix}${var.test_suffix}"
  backend_service = google_compute_backend_service.gcp_compute_target_ssl_proxies_backend_service.id
  ssl_certificates = [
  google_compute_ssl_certificate.gcp_compute_ssl_certificates_cert.id]

}

resource "google_compute_backend_service" "gcp_compute_target_ssl_proxies_backend_service" {
  name     = "ssl-service-${var.test_prefix}${var.test_suffix}"
  protocol = "SSL"
  health_checks = [
  google_compute_health_check.gcp_compute_target_ssl_proxies_health_check.id]
}

resource "google_compute_health_check" "gcp_compute_target_ssl_proxies_health_check" {
  name               = "ssl-health-check-${var.test_prefix}${var.test_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "443"
  }
}