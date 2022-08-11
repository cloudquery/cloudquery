terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "ecs"
    region         = "us-east-1"
  }
}
