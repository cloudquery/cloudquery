terraform {
  backend "s3" {
    bucket = "cq-provider-aws-tf"
    key    = "eventbridge"
    region = "us-east-1"
  }
}
