terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "appsync"
    region = "us-east-1"
  }
}
