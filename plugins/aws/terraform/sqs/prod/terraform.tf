terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "sqs"
    region         = "us-east-1"
  }
}
