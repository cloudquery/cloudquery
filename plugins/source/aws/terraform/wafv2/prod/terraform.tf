terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "wafv2"
    region = "us-east-1"
  }
}
