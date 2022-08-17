terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "ssm"
    region = "us-east-1"
  }
}
