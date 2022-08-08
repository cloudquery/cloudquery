terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "glue"
    region         = "us-east-1"
  }
}
