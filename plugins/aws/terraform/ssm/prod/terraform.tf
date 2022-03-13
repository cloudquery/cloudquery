terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "ssm"
    region         = "us-east-1"
  }
}
