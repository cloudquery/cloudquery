terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "iot"
    region         = "us-east-1"
  }
}
