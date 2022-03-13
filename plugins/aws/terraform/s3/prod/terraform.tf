terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "s3"
    region         = "us-east-1"
  }
}
