cloudquery {
  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    source  = "cloudquery"
    version = "v0.0.0"
  }
}

provider "aws" {
  configuration {
    account "dev" {
      role_arn = "12312312"
    }
    account "ron" {}
  }
  resources = ["slow_resource"]
}

provider "aws" {
  alias = "another-aws"
  configuration {
    account "dev" {
      role_arn = "12312312"
    }
    account "ron" {}
  }
  resources = ["slow_resource"]
}