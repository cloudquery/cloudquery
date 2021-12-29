terraform {
  backend "s3" {
    bucket = "cq-provider-aws-tf-state"
    key    = "tf-state-perm"
    region = "eu-central-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.55"
    }
  }
}