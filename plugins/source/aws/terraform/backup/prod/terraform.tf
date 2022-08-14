terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "backup"
    region = "us-east-1"
  }
}
