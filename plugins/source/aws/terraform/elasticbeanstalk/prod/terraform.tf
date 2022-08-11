terraform {
  backend "s3" {
    bucket = "cq-provider-aws-tf"
    key    = "elasticbeanstalk"
    region = "us-east-1"
  }
}
