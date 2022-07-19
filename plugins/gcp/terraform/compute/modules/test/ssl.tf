################################################################################
# Compute SSL Module - Policies
################################################################################

resource "google_compute_ssl_policy" "ssl-policy" {
  name    = "${var.prefix}-ssl-policy"
  profile = "MODERN"
}

################################################################################
# Compute SSL Module - Certificate
################################################################################

resource "google_compute_managed_ssl_certificate" "gcp_compute_managed_ssl_certificates_cert" {
  name = "${var.prefix}-managed-cert"

  managed {
    domains = [var.domain, "ex.${var.domain}"]
  }
}

#####################