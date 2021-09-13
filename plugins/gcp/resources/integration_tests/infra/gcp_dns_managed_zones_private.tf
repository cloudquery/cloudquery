resource "google_dns_managed_zone" "gcp_dns_managed_zones_zone_private" {
  name        = "private-zone${var.test_prefix}${var.test_suffix}"
  dns_name    = "example-p-${var.test_suffix}.com."
  description = "Example DNS zone"
  labels = {
    test = "test"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network.id
    }
    networks {
      network_url = google_compute_network.network1.id
    }
  }

  forwarding_config {
    target_name_servers {
      ipv4_address = "172.16.1.10"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }
}
