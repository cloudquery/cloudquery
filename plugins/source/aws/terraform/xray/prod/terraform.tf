terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "xray"
    region         = "us-east-1"
  }
}
