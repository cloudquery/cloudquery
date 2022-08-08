terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "backup"
    region         = "us-east-1"
  }
}
