terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "athena"
    region         = "us-east-1"
  }
}
