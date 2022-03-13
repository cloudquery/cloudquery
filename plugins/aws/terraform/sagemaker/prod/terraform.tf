terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "sagemaker"
    region         = "us-east-1"
  }
}
