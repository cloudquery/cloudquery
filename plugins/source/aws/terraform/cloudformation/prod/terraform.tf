terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "cloudformation"
    region         = "us-east-1"
  }
}
