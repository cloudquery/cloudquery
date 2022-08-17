terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "sagemaker"
    region = "us-east-1"
  }
}
