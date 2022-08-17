terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "codepipeline"
    region = "us-east-1"
  }
}
