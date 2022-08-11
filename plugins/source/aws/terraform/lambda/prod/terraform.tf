terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "lambda"
    region         = "us-east-1"
  }
}
