provider "aws" {
  region  = "us-east-1"

  default_tags {
    tags = {
      TestId = var.test_suffix
      Type = "integration_test"
    }
  }
}