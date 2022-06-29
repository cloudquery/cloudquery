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
    accounts "dev" {
      role_arn = "${ROLE_ARN}"
    }
    accounts "ron" {}
  }
  resources = ["slow_resource"]
}