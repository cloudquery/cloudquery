terraform {
  backend "s3" {
    bucket = "cq-provider-aws-tf-state"
    key    = "tf-state2"
    region = "eu-central-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "= 3.66.0"
    }
  }
}