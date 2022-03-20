module "vpc" {
  source  = "terraform-google-modules/network/google"
  version = "~> 4.0.0"

  project_id   = var.project_id
  network_name = "${var.prefix}-compute-vpc"
  routing_mode = "GLOBAL"
  description  = "Private network for cq-provider-gcp/compute"

  subnets = [
    {
      subnet_name   = "${var.prefix}-compute-subnet-01"
      subnet_ip     = "10.10.20.0/24"
      subnet_region = "${var.region}"
    }
  ]
}