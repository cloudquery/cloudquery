terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "ses"
    region = "us-east-1"
  }
}
