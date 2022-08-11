terraform {
  backend "s3" {
    bucket         = "cq-provider-cf-tf"
    key            = "account"
    region         = "us-east-1"
  }
}
