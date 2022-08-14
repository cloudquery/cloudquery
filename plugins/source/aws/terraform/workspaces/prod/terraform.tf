terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "workspaces"
    region = "us-east-1"
  }
}
