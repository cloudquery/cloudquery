terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "sqs"
    region = "us-east-1"
  }
}
