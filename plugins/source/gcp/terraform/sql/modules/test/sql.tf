# https://github.com/terraform-google-modules/terraform-google-sql-db/blob/master/examples/postgresql-public/main.tf

# module "network" {
#   source  = "terraform-google-modules/network/google"
#   version = "~> 4.0.0"

#   project_id   = var.project_id
#   network_name = "${var.prefix}-sql"
#   routing_mode = "GLOBAL"
#   description  = "Private network for cq-provider-gcp/sql"

#   subnets = [
#     {
#       subnet_name   = "subnet-01"
#       subnet_ip     = "10.10.20.0/24"
#       subnet_region = "${var.region}-a"
#     }
#   ]
# }

module "postgresql-db" {
  source  = "GoogleCloudPlatform/sql-db/google//modules/postgresql"
  version = "10.0.0"
  name                 = "${var.prefix}-sql-pgsql"
  random_instance_name = true
  database_version     = "POSTGRES_9_6"
  project_id           = var.project_id
  zone                 = "${var.region}-b"
  region               = var.region
  tier                 = "db-f1-micro"

  deletion_protection = false
  create_timeout = "25m"
  ip_configuration = {
    ipv4_enabled        = true
    private_network     = null
    require_ssl         = true
    allocated_ip_range  = null
    authorized_networks = []
  }
}