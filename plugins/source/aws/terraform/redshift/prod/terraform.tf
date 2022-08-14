terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "redshift"
    region = "us-east-1"
  }
}
