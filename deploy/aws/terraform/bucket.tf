resource "aws_s3_bucket" "deploy_bucket" {
  bucket = var.bucket
  acl    = "private"

  tags = {
    Name        = var.bucket
  }
}