terraform {
  backend "s3" {
    bucket = "cq-provider-aws-tf-state"
    key    = "tf-state2"
    region = "eu-central-1"
    dynamodb_table = "terraform-lock"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
  }
}