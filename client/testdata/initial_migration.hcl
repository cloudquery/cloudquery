cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "aws" {
    source = "cloudquery/cq-provider-aws"
    version = "v0.5.7"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "aws" {
  configuration {}
  resources = [
    "*"]
}