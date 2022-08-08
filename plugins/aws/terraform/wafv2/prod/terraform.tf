terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "wafv2"
    region         = "us-east-1"
  }
}
