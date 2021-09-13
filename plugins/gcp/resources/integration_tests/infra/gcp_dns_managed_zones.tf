
resource "google_dns_managed_zone" "gcp_dns_managed_zones_zone_public" {
  name        = "managed-zone${var.test_prefix}${var.test_suffix}"
  dns_name    = "example-${var.test_suffix}.com."
  description = "Example DNS zone"
  labels = {
    test = "test"
  }

  visibility = "public"

  dnssec_config {
    default_key_specs {
      algorithm  = "rsasha256"
      key_length = 2048
      key_type   = "keySigning"
      kind       = "dns#dnsKeySpec"
    }

    default_key_specs {
      algorithm  = "rsasha256"
      key_length = 1024
      key_type   = "zoneSigning"
      kind       = "dns#dnsKeySpec"
    }
  }
}

