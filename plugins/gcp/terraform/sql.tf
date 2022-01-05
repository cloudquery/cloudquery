################################################################################
# SQL Helper
################################################################################

locals {
  ip_configuration = [
    {
      ipv4_enabled        = true
      private_network     = null
      require_ssl         = false
      authorized_networks = []
    }
  ]
}

resource "google_sql_database_instance" "mysql-replica" {
  provider             = google-beta
  project              = local.project
  name                 = "${module.private-mysql-db.instance_name}-replica"
  database_version     = "MYSQL_5_6"
  region               = local.region
  master_instance_name = module.private-mysql-db.instance_name
  deletion_protection  = false

  replica_configuration {
    failover_target = false
  }

  settings {
    tier              = "db-n1-standard-1"
    activation_policy = "ALWAYS"

    dynamic "ip_configuration" {
      for_each = local.ip_configuration
      content {
        ipv4_enabled    = lookup(ip_configuration.value, "ipv4_enabled", null)
        private_network = lookup(ip_configuration.value, "private_network", null)
        require_ssl     = lookup(ip_configuration.value, "require_ssl", null)

        dynamic "authorized_networks" {
          for_each = lookup(ip_configuration.value, "authorized_networks", [])
          content {
            expiration_time = lookup(authorized_networks.value, "expiration_time", null)
            name            = lookup(authorized_networks.value, "name", null)
            value           = lookup(authorized_networks.value, "value", null)
          }
        }
      }
    }

    disk_autoresize = true
    disk_size       = 10
    disk_type       = "PD_SSD"
    pricing_plan    = "PER_USE"
    user_labels     = local.labels
  }
  depends_on = [module.private-mysql-db]
}

################################################################################
# SQL Module
################################################################################

module "mysql-private-service-access" {
  source      = "GoogleCloudPlatform/sql-db/google//modules/private_service_access"
  version     = "~> 8.0.0"
  project_id  = local.project
  vpc_network = module.vpc.network_name
}

module "private-mysql-db" {
  source               = "GoogleCloudPlatform/sql-db/google//modules/mysql"
  version              = "~> 8.0.0"
  name                 = "${local.prefix}-private-mysql-db"
  random_instance_name = true
  project_id           = local.project

  zone = "${local.region}-a"

  disk_type = "PD_SSD"
  disk_size = 10

  deletion_protection = false

  database_version = "MYSQL_5_6"
  region           = local.region
  tier             = "db-n1-standard-1"

  module_depends_on = [module.mysql-private-service-access.peering_completed]

  user_labels = local.labels
}