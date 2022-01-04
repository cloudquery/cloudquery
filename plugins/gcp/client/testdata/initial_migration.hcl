cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "gcp" {
    source = "cloudquery/cq-provider-gcp"
    version = "v0.5.1"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "gcp" {
  configuration {}
  resources = [
    "*"]
}
