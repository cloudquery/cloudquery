terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "autoscaling"
    region = "us-east-1"
  }
}
