terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "lambda"
    region = "us-east-1"
  }
}
