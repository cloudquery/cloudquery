terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "accessanalyzer"
    region         = "us-east-1"
  }
}
