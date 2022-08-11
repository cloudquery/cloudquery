terraform {
  backend "s3" {
    bucket         = "cq-provider-aws-tf"
    key            = "workspaces"
    region         = "us-east-1"
  }
}
