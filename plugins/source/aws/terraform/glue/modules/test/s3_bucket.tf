resource "aws_s3_bucket" "aws_s3_bucket" {
  bucket        = "${var.prefix}-glue-target-bucket"
  force_destroy = true
}