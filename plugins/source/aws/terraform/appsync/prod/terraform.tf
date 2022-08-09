terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "appsync"
    region         = "us-east-1"
  }
}
