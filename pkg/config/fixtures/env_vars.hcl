cloudquery {
  connection {
    dsn = "${DSN}"
  }
  provider "test" {
    source  = "cloudquery"
    version = "v0.0.0"
  }
}

provider "aws" {
  configuration {
    account "dev" {
      role_arn = "${ROLE_ARN}"
    }
    account "ron" {}
  }
  resources = ["slow_resource"]
}