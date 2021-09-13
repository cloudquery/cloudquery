resource "google_compute_target_https_proxy" "gcp_compute_target_https_proxies" {
  name    = "https-proxy-${var.test_prefix}${var.test_suffix}"
  url_map = google_compute_url_map.gcp_compute_target_https_proxies_urlmap.id
  ssl_certificates = [
  google_compute_ssl_certificate.gcp_compute_ssl_certificates_cert.id]
  ssl_policy = google_compute_ssl_policy.custom-ssl-policy.id
}

resource "google_compute_url_map" "gcp_compute_target_https_proxies_urlmap" {
  name        = "https-urlmap${var.test_prefix}${var.test_suffix}"
  description = "a description"

  default_service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id

  host_rule {
    hosts = [
    "mysite.com"]
    path_matcher = "mysite"
  }

  host_rule {
    hosts = [
    "myothersite.com"]
    path_matcher = "otherpaths"
  }

  path_matcher {
    name            = "mysite"
    default_service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id

    path_rule {
      paths = [
      "/home"]
      service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id
    }

    path_rule {
      paths = [
      "/static"]
      service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id
    }
  }

  path_matcher {
    name            = "otherpaths"
    default_service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id
  }

  test {
    service = google_compute_backend_bucket.gcp_compute_target_https_proxies_compute_backend_bucket.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_bucket" "gcp_compute_target_https_proxies_compute_backend_bucket" {
  name        = "https-backend-${var.test_prefix}${var.test_suffix}"
  bucket_name = google_storage_bucket.gcp_compute_target_https_proxies_storage_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "gcp_compute_target_https_proxies_storage_bucket" {
  name     = "https-bucket-${var.test_prefix}${var.test_suffix}"
  location = "US"
}
