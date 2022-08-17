terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "glue"
    region = "us-east-1"
  }
}
