cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "azure" {
    source = "cloudquery/cq-provider-azure"
    version = "v0.3.7"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "azure" {
  configuration {}
  resources = [
    "*"]
}