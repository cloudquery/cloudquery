cloudquery {

  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    source = "cloudquery"
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