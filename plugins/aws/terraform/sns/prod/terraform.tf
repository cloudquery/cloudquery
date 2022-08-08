terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "sns"
    region         = "us-east-1"
  }
}
