resource "google_compute_interconnect_attachment" "gcp_compute_interconnects_interconnect" {
  name                     = "on-prem-attachment"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  type                     = "PARTNER"
  router                   = google_compute_router.gcp_compute_interconnects_router.id
  mtu                      = 1500
}

resource "google_compute_router" "gcp_compute_interconnects_router" {
  name    = "router"
  network = google_compute_network.network.name
  bgp {
    asn = 16550
  }
}
