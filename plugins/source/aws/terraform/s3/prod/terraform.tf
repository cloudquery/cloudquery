terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "s3"
    region = "us-east-1"
  }
}
