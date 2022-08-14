terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "mq"
    region = "us-east-1"
  }
}
