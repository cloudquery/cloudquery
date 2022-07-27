resource "aws_glue_dev_endpoint" "aws_glue_dev_endpoint" {
  name     = "${var.prefix}-dev-endpoint"
  role_arn = aws_iam_role.aws_iam_role.arn
}

resource "aws_iam_role" "aws_iam_role" {
  name               = "AWSGlueServiceRole-${var.prefix}-role"
  assume_role_policy = data.aws_iam_policy_document.aws_iam_policy_document.json
}

data "aws_iam_policy_document" "aws_iam_policy_document" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["glue.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "aws_iam_role_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"
  role       = aws_iam_role.aws_iam_role.name
}