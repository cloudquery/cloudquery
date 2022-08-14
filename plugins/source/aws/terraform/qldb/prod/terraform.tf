terraform {
  backend "s3" {
    bucket = "cq-plugins-source-aws-tf"
    key    = "qldb"
    region = "us-east-1"
  }
}
