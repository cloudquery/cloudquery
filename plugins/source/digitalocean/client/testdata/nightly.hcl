cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "digitalocean" {
    source  = "cloudquery/cq-provider-digitalocean"
    version = "latest"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "digitalocean" {
  configuration {
    // token = FROM ENV
     spaces_regions = ["nyc3", "sfo3", "ams3", "sgp1", "fra1"]
    // spaces_access_key = FROM ENV
    // spaces_access_key_id = FROM ENV
     spaces_debug_logging = false
  }
  resources = ["*"]
}