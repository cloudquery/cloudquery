cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "aws" {
    source = "cloudquery/cq-provider-aws"
    version = "latest"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "aws" {
  configuration {

    accounts "default" {
      role_arn = "${ASSUME_ROLE_ARN}"
    }

    max_retries = 20
    max_backoff = 60
  }
  resources = [
    "*"]
}
