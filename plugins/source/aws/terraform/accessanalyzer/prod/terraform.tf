terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws"
    key    = "accessanalyzer"
    region = "us-east-1"
  }
}
