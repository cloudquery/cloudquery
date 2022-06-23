terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "lightsail"
    region         = "us-east-1"
  }
}
