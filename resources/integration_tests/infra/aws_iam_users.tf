resource "aws_iam_user" "iam_user" {
  name = "user${var.test_prefix}${var.test_suffix}"
  path = "/system/"

  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_access_key" "iam_user_acc_key" {
  user = aws_iam_user.iam_user.name
}

resource "aws_iam_user_policy" "user_inline_policy" {
  name = "user_policy${var.test_prefix}${var.test_suffix}"
  user = aws_iam_user.iam_user.name

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

resource "aws_iam_policy" "user_policy" {
  name        = "policy${var.test_prefix}${var.test_suffix}"
  description = "A test policy"
  policy      = <<EOF
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

resource "aws_iam_user_policy_attachment" "user_policy_attach" {
  user       = aws_iam_user.iam_user.name
  policy_arn = aws_iam_policy.user_policy.arn
}

resource "aws_iam_group_membership" "team" {
  name = "membership${var.test_prefix}${var.test_suffix}"

  users = [
    aws_iam_user.iam_user.name,
  ]

  group = aws_iam_group.group_developers.name
}


