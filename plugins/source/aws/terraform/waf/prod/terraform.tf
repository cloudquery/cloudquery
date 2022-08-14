terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "waf"
    region = "us-east-1"
  }
}
