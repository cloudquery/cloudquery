terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "codepipeline"
    region         = "us-east-1"
  }
}
