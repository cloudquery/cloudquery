################################################################################
# Compute SSL Module - Policies
################################################################################

resource "google_compute_ssl_policy" "ssl-policy" {
  name    = "${local.prefix}-ssl-policy"
  profile = "MODERN"
}

################################################################################
# Compute SSL Module - Certificate
################################################################################

resource "google_compute_managed_ssl_certificate" "gcp_compute_managed_ssl_certificates_cert" {
  name = "${local.prefix}-managed-cert"

  managed {
    domains = [local.domain]
  }
}