cloudquery {
  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    version = "v0.0.0"
  }
}

provider "test" {
  configuration {
    accounts "dev" {
      role_arn = "12312312"
    }
    accounts "ron" {}
  }
  resources = ["slow_resource"]
}