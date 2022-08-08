terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "autoscaling"
    region         = "us-east-1"
  }
}
