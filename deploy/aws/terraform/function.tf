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

data "archive_file" "zip" {
  type        = "zip"
  source_file = "../../../bin/cloudquery"
  output_path = "cloudquery.zip"
}


resource "aws_iam_role_policy_attachment" "cloudquery_role_attachment" {
  role       = aws_iam_role.iam_for_cloudquery.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}


resource "aws_lambda_function" "cloudquery" {
  handler       = "cloudquery"
  function_name = "cloudquery"
  filename      = "cloudquery.zip"
  runtime       = "go1.x"
  role          = aws_iam_role.iam_for_cloudquery.arn

  source_code_hash = data.archive_file.zip.output_base64sha256

  environment {
    variables = {
      CQ_DRIVER= "mysql",
      CQ_DSN = "${aws_rds_cluster.cloudquery.master_username}:${aws_rds_cluster.cloudquery.master_password}@tcp(${aws_rds_cluster.cloudquery.endpoint}:3306)/cloudquery"
    }
  }
}