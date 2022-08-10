resource "aws_glue_job" "example" {
  name     = "${var.prefix}-glue-job"
  role_arn = aws_iam_role.aws_iam_role.arn

  command {
    script_location = "s3://${aws_s3_bucket.aws_s3_bucket.bucket}/scripts/example.py"
  }
}