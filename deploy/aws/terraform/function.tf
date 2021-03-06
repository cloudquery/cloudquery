resource "aws_iam_role" "iam_for_cloudquery" {
  name = "cloudquery_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "policy_document" {
  statement {
    sid = "1"
    actions = [
      "s3:GetObject"
    ]
    resources = [
      "arn:aws:s3:::${var.bucket}"
    ]
  }
}

resource "aws_iam_policy" "policy" {
  name = "cloudquery_policy"
  path = "/"
  policy = data.aws_iam_policy_document.policy_document.json
}

resource "aws_s3_bucket_object" "file_upload" {
  bucket = var.bucket
  key    = "lambda-functions/cloudquery.zip"
  source = data.archive_file.zip.output_path

  etag   = filemd5(data.archive_file.zip.output_path)

  depends_on = [aws_s3_bucket.deploy_bucket]

}

data "archive_file" "zip" {
  type        = "zip"
  source_file = "../../../bin/cloudquery"
  output_path = "cloudquery.zip"
}


resource "aws_iam_role_policy_attachment" "cloudquery_role_attachment1" {
  role       = aws_iam_role.iam_for_cloudquery.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

resource "aws_iam_role_policy_attachment" "cloudquery_role_attachment2" {
  role       = aws_iam_role.iam_for_cloudquery.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

resource "aws_iam_role_policy_attachment" "cloudquery_role_attachment3" {
  role       = aws_iam_role.iam_for_cloudquery.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_lambda_function" "cloudquery" {
  handler       = "cloudquery"
  function_name = "cloudquery"
  s3_bucket = var.bucket
  s3_key = "lambda-functions/cloudquery.zip"
  runtime       = "go1.x"
  role          = aws_iam_role.iam_for_cloudquery.arn
  timeout       = 900
  memory_size   = 256

  source_code_hash = data.archive_file.zip.output_base64sha256

  depends_on = [
    aws_s3_bucket_object.file_upload
  ]

  vpc_config {
    subnet_ids         = [aws_subnet.rds_subnet_a.id, aws_subnet.rds_subnet_b.id]
    security_group_ids = [aws_security_group.allow_postgresql.id, aws_security_group.allow_egress.id]
  }

  environment {
    variables = {
      CQ_DRIVER= "postgresql",
      CQ_DSN = "user=${aws_rds_cluster.cloudquery.master_username} password=${aws_rds_cluster.cloudquery.master_password} host=${aws_rds_cluster.cloudquery.endpoint} port=5432 dbname=cloudquery",
      CQ_PLUGIN_DIR = "/tmp"
    }
  }
}