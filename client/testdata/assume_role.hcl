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
      account_id = "${ASSUME_ROLE_ACCOUNT_ID}"
    }

    max_retries = 20
    max_backoff = 60
  }
  resources = [
    "*"]
}
