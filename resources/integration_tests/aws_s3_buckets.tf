resource "aws_s3_bucket" "b" {
  bucket = "${var.test_prefix}${var.test_suffix}"
  acl = "private"
}