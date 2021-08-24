resource "aws_s3_bucket" "s3_bucket" {
  bucket = "bucket-${var.test_prefix}${var.test_suffix}"
  acl = "private"
}