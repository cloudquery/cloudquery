terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "sns"
    region = "us-east-1"
  }
}
