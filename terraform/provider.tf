provider "aws" {
  region = "eu-central-1"

  default_tags {
    tags = {
      Type = "cq-provider-aws-test"
    }
  }
}

// data "aws_region" "current" {}