terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "waf"
    region         = "us-east-1"
  }
}
