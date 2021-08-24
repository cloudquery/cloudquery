resource "aws_iam_group" "group_developers" {
  name = "aws_iam_group${var.test_prefix}${var.test_suffix}"
  path = "/users/"
}

resource "aws_iam_group_policy" "group_policy" {
  name = "aws_iam_group_policy${var.test_prefix}${var.test_suffix}"
  group = aws_iam_group.group_developers.name

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

// TODO add group policy attachment