terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "elasticbeanstalk"
    region = "us-east-1"
  }
}
