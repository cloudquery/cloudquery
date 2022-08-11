terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "rds"
    region         = "us-east-1"
  }
}
