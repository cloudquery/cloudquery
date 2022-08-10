terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "mq"
    region         = "us-east-1"
  }
}
