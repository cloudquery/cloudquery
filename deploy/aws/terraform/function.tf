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


resource "aws_iam_role_policy_attachment" "cloudquery_role_attachment" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}


resource "aws_lambda_function" "cloudquery" {
  handler       = "cloudquery"
  function_name = "cloudquery"
  filename      = "cloudquery.zip"
  runtime       = "go1.x"
  role          = aws_iam_role.iam_for_lambda.arn

  environment {
    variables = {
      CLOUDQUERY_DRIVER = "mysql",
      CLOUDQUERY_DATABASE_STRING = "${aws_rds_cluster.cloudquery.master_username}:${aws_rds_cluster.cloudquery.master_password}@tcp(${aws_rds_cluster.cloudquery.endpoint}:3306)/${aws_rds_cluster.cloudquery.database_name}"
    }
  }
}