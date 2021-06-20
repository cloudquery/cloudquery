cloudquery {

  connection {
    dsn = "host=localhost user=postgres password=pass DB.name=postgres port=5432"
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