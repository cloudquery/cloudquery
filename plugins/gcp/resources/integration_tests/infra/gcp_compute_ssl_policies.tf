resource "google_compute_ssl_policy" "prod-ssl-policy" {
  name    = "prod-ssl-policy-${var.test_prefix}${var.test_suffix}"
  profile = "MODERN"
}

resource "google_compute_ssl_policy" "nonprod-ssl-policy" {
  name            = "nonprod-ssl-policy-${var.test_prefix}${var.test_suffix}"
  profile         = "MODERN"
  min_tls_version = "TLS_1_2"
}

resource "google_compute_ssl_policy" "custom-ssl-policy" {
  name            = "custom-ssl-policy-${var.test_prefix}${var.test_suffix}"
  min_tls_version = "TLS_1_2"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}

