terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "cloudformation"
    region = "us-east-1"
  }
}
