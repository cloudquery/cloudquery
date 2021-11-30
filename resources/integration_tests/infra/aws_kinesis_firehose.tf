resource "aws_kinesis_firehose_delivery_stream" "kinesis_firehose_delivery_stream" {
  name        = "aws-waf-logs-kinesis-firehose-${var.test_prefix}${var.test_suffix}"
  destination = "extended_s3"

  extended_s3_configuration {
    role_arn   = aws_iam_role.firehose_delivery_stream_bucket_role.arn
    bucket_arn = aws_s3_bucket.kinesis_firehose_delivery_stream_bucket.arn
  }
}

resource "aws_s3_bucket" "kinesis_firehose_delivery_stream_bucket" {
  bucket = "kinesis-firehose-bucket-${var.test_prefix}${var.test_suffix}"
  acl    = "private"
}

resource "aws_iam_role" "firehose_delivery_stream_bucket_role" {
  name = "kinesis-firehose-role-${var.test_prefix}${var.test_suffix}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}
