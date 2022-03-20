################################################################################
# Network VPN - Helpers
################################################################################
#
#resource "google_compute_interconnect_attachment" "gcp_compute_interconnects_interconnect" {
#  name                     = "${local.prefix}-ico-attachment"
#  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
#  type                     = "PARTNER"
#  router                   = google_compute_router.interconnect_router.id
#  mtu                      = 1500
#}
#
#resource "google_compute_router" "interconnect_router" {
#  name    = "${local.prefix}-${local.region}-interconnect-router"
#  network = module.vpc.network_name
#  bgp {
#    asn = 16550
#  }
#}

resource "google_compute_router" "vpn_router" {
  name    = "${var.prefix}-${var.region}-compute-vpn-router"
  region  = var.region
  network = module.vpc.network_name
  project = var.project_id

  bgp {
    asn = "64519"
  }
}

################################################################################
# Network VPN - Helpers
################################################################################

module "vpn-manage-internal" {
  source  = "terraform-google-modules/vpn/google"
  version = "~> 1.2.0"
  project_id         = var.project_id
  network            = module.vpc.network_id
  region             = var.region
  gateway_name       = "${var.prefix}-compute"
  tunnel_name_prefix = "${var.prefix}-compute-tunnel"
  # shared_secret      = "secrets"
  tunnel_count       = 1
  peer_ips           = ["1.1.1.1", "2.2.2.2"]

  route_priority = 1000
  remote_subnet  = ["10.10.12.0/24"]
}