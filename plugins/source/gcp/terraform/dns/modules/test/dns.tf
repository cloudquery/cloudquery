################################################################################
# Dns Module - Helpers
################################################################################

# resource "google_dns_policy" "dns-policy" {
#   name                      = "${local.prefix}-dns-policy"
#   enable_inbound_forwarding = true

#   enable_logging = true

#   alternative_name_server_config {
#     target_name_servers {
#       ipv4_address    = "10.10.10.31"
#       forwarding_path = "private"
#     }
#   }

#   networks {
#     network_url = module.vpc.network_id
#   }
# }

################################################################################
# Dns Module - Forwarding
################################################################################

# module "dns-forwarding-zone" {
#   source     = "terraform-google-modules/cloud-dns/google"
#   version    = "3.0.0"
#   project_id = local.project
#   type       = "forwarding"
#   name       = "${local.prefix}-forwarding-zone"
#   domain     = "${local.prefix}-forwarding-zone.com."

#   private_visibility_config_networks = [
#     module.vpc.network_self_link
#   ]

#   target_name_server_addresses = ["10.10.10.44"]
# }

################################################################################
# Dns Module - Private
################################################################################

module "dns-private-zone" {
  source     = "terraform-google-modules/cloud-dns/google"
  version    = "4.1.0"
  project_id = var.project_id
  type       = "private"
  name       = "${var.prefix}-dns-private-zone"
  domain     = "${var.prefix}.cq-provider-gcp.cloudquery.io."

  private_visibility_config_networks = [
    module.vpc.network_self_link
  ]

  recordsets = [
    {
      name    = ""
      type    = "NS"
      ttl     = 300
      records = [
        "ns.${var.prefix}.cq-provider-gcp.cloudquery.io.",
      ]
    },
    {
      name    = "localhost"
      type    = "A"
      ttl     = 300
      records = [
        "10.10.10.31",
      ]
    },
  ]

  labels = var.labels
}

################################################################################
# Dns Module - Public
################################################################################

# module "dns-public-zone" {
#   source     = "terraform-google-modules/cloud-dns/google"
#   version    = "4.1.0"
#   project_id = local.project
#   type       = "public"
#   name       = "${local.prefix}-public-zone"
#   domain     = "${local.prefix}-public-zone.com."

#   private_visibility_config_networks = [
#     module.vpc.network_self_link
#   ]

#   recordsets = [
#     {
#       name    = ""
#       type    = "NS"
#       ttl     = 300
#       records = [
#         "ns.${local.prefix}-public-zone.com.",
#       ]
#     },
#     {
#       name    = "localhost"
#       type    = "A"
#       ttl     = 300
#       records = [
#         "10.10.10.30",
#       ]
#     },
#   ]

#   dnssec_config = {
#     algorithm  = "rsasha256"
#     key_length = 1024
#     key_type   = "zoneSigning"
#     kind       = "dns#dnsKeySpec"
#   }

#   labels = local.labels
# }