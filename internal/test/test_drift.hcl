cloudquery {

  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
  }
  provider "aws" {
    version = "latest"
  }

}

modules {
  // drift configuration block
  drift "s3" {
    terraform {
      backend = "s3"
      bucket = "cq-provider-aws-tf"
      keys = [ "*" ]
    }
    provider "aws" {
      account_ids     = ["1234567891011"]
    }
  }
}