cloudquery {
  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    source  = "cloudquery"
    version = "0.0.0"
  }
}