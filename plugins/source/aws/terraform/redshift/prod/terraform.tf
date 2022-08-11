terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "redshift"
    region         = "us-east-1"
  }
}
