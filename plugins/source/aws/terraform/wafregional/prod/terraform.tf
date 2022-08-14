terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "wafregional"
    region = "us-east-1"
  }
}
