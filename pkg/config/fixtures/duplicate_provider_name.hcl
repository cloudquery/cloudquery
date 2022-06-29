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
    accounts "dev" {
      role_arn = "12312312"
    }
    accounts "ron" {}
  }
  resources = ["slow_resource"]
}

provider "aws" {
  configuration {
    accounts "dev" {
      role_arn = "12312312"
    }
    accounts "ron" {}
  }
  resources = ["slow_resource"]
}