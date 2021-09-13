resource "google_sql_database_instance" "sql_db_instance" {
  name                = "sql-database-instance-${var.test_prefix}${var.test_suffix}"
  database_version    = "POSTGRES_11"
  deletion_protection = false

  settings {
    tier = "db-f1-micro"

    ip_configuration {
      authorized_networks {
        name  = "testnet"
        value = "8.8.8.8"
      }
    }
  }
}
