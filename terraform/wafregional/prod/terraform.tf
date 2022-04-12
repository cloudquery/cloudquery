terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "wafregional"
    region         = "us-east-1"
  }
}
