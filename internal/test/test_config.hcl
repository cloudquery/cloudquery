cloudquery {

  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
  }
  provider "test" {
    version = "latest"
  }

}

// All Provider Configurations
provider "test" {
  configuration {}

  resources = [
    "slow_resource"
  ]
}