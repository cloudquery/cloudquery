terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "ecs"
    region = "us-east-1"
  }
}
