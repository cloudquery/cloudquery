terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "ses"
    region         = "us-east-1"
  }
}
