terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "qldb"
    region         = "us-east-1"
  }
}
