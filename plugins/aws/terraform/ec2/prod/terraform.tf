terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "ec2"
    region         = "us-east-1"
  }
}
