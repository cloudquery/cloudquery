resource "google_compute_network" "network" {
  name = "network-${var.test_prefix}${var.test_suffix}"
  mtu  = 1500
}

resource "google_compute_network" "network1" {
  name                    = "network1-${var.test_prefix}${var.test_suffix}"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "network-subnetwork" {
  name          = "network-subnetwork-${var.test_prefix}${var.test_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  network       = google_compute_network.network.id
  region        = var.region
  secondary_ip_range {
    range_name    = "range-${var.test_prefix}${var.test_suffix}"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network_peering" "network-peering" {
  name         = "network-peering-${var.test_prefix}${var.test_suffix}"
  network      = google_compute_network.network.id
  peer_network = google_compute_network.network1.id
}

resource "google_compute_vpn_gateway" "vpn_gateway" {
  name        = "vpn-gateway-${var.test_prefix}${var.test_suffix}"
  network     = google_compute_network.network.id
  description = "a description"
}



