terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "xray"
    region = "us-east-1"
  }
}
