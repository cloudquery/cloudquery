cloudquery {

  connection {
    dsn = "tsdb://postgres:pass@localhost:5432/postgres?sslmode=disable"
  }
  provider "test" {
    source = "cloudquery"
    version = "latest"
  }
  history {
    // Save data retention for 7 days
    retention = 7
    // Truncate our fetch by 6 hours per fetch
    truncation = 6
  }

}

// All Provider Configurations
provider "test" {
  configuration {}

  resources = [
    "slow_resource"
  ]
}