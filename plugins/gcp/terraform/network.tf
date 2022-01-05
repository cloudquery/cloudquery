################################################################################
# Network Module
################################################################################

module "vpc" {
  source  = "terraform-google-modules/network/google"
  version = "~> 4.0.0"

  project_id   = local.project
  network_name = "${local.prefix}-private-network"
  routing_mode = "GLOBAL"
  description  = "Private network for CQ"

  subnets = [
    {
      subnet_name   = "${local.prefix}-private-network-subnet-01"
      subnet_ip     = "${local.subnet_ids.0}/24"
      subnet_region = local.region
    },
    {
      subnet_name   = "${local.prefix}-network-subnet-02"
      subnet_ip     = "${local.subnet_ids.1}/24"
      subnet_region = local.region
    },
    {
      subnet_name   = "${local.prefix}-network-subnet-03"
      subnet_ip     = "${local.subnet_ids.2}/24"
      subnet_region = local.region
    }
  ]

  routes = [
    {
      name              = "${local.prefix}-egress-internet"
      description       = "route through IGW to access internet"
      destination_range = "0.0.0.0/0"
      tags              = "egress-inet"
      next_hop_internet = "true"
    }
  ]

  firewall_rules = [
    {
      name        = "${local.prefix}-compute-firewall-a"
      network     = module.vpc.network_name
      direction   = "INGRESS"
      allow       = [
        {
          protocol = "tcp"
          ports    = [
            "80",
            "443"
          ]
        }
      ]
      source_tags = [
        "web"
      ]
      target_tags = local.tags
    },
    {
      name        = "${local.prefix}-compute-firewall-b"
      network     = module.vpc.network_name
      direction   = "INGRESS"
      deny        = [
        {
          protocol = "tcp"
          ports    = [
            "22",
          ]
        }
      ]
      source_tags = [
        "ssh"
      ]
      target_tags = local.tags
    }
  ]

  secondary_ranges = {
    "${local.prefix}-private-network-subnet-01" = [
      {
        range_name    = "${local.prefix}-private-network-secondary-01"
        ip_cidr_range = "192.168.64.0/24"
      }
    ]
  }
}

module "another-network" {
  source  = "terraform-google-modules/network/google"
  version = "~> 4.0.0"

  project_id   = local.project
  network_name = "${local.prefix}-another-network"
  routing_mode = "GLOBAL"
  description  = "Private network for CQ"

  subnets = [
    {
      subnet_name   = "${local.prefix}-another-network-subnet-01"
      subnet_ip     = "${local.subnet_ids.0}/24"
      subnet_region = local.region
    }
  ]
}