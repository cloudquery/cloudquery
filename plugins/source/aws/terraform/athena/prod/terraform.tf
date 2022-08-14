terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "athena"
    region = "us-east-1"
  }
}
