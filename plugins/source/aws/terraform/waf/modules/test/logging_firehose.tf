resource "aws_kinesis_firehose_delivery_stream" "extended_s3_stream" {
  name        = "aws-waf-logs-${var.prefix}-logs-kinesis-firehose"
  destination = "extended_s3"

  extended_s3_configuration {
    role_arn   = aws_iam_role.firehose_role.arn
    bucket_arn = aws_s3_bucket.bucket.arn

    processing_configuration {
      enabled = "true"

      processors {
        type = "Lambda"

        parameters {
          parameter_name  = "LambdaArn"
          parameter_value = "${aws_lambda_function.lambda_processor.arn}:$LATEST"
        }
      }
    }
  }

  tags = merge(
    var.tags,
    {
      Name = "aws-waf-logs-${var.prefix}-logs-kinesis-firehose"
    }
  )
}

resource "aws_s3_bucket" "bucket" {
  bucket = "${var.prefix}-waf-firehose-logging-bucket"

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-waf-firehose-logging-bucket"
    }
  )
}

resource "aws_s3_bucket_acl" "bucket_acl" {
  bucket = aws_s3_bucket.bucket.id
  acl    = "private"
}

resource "aws_iam_role" "firehose_role" {
  name = "${var.prefix}-waf-firehose-role"

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

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-waf-firehose-role"
    }
  )
}

resource "aws_iam_role" "lambda_role" {
  name = "${var.prefix}-waf-firehose-s3-lambda-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-waf-firehose-s3-lambda-role"
    }
  )
}

data "archive_file" "processor_package" {
  type        = "zip"
  source_file = "${path.root}/../modules/test/fixtures/processor.py"
  output_path = "${path.root}/../modules/test/fixtures/processor.zip"
}

resource "aws_lambda_function" "lambda_processor" {
  filename         = "${path.root}/../modules/test/fixtures/processor.zip"
  function_name    = "${var.prefix}-waf-firehose-s3-lambda-func"
  role             = aws_iam_role.lambda_role.arn
  runtime          = "python3.8"
  source_code_hash = data.archive_file.processor_package.output_base64sha256
  handler          = "processor.lambda_handler"
  timeout          = 10

  tags = merge(
    var.tags,
    {
      "Name" = "${var.prefix}-waf-firehose-s3-lambda-func"
    }
  )

}