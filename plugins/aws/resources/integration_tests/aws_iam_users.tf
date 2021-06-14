resource "aws_iam_user" "lb" {
  name = "user${var.test_prefix}${var.test_suffix}"
  path = "/system/"


  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_access_key" "lb" {
  user = aws_iam_user.lb.name
}

resource "aws_iam_user_policy" "lb_ro" {
  name = "user_policy${var.test_prefix}${var.test_suffix}"
  user = aws_iam_user.lb.name

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

resource "aws_iam_policy" "policy" {
  name = "policy${var.test_prefix}${var.test_suffix}"
  description = "A test policy"
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

resource "aws_iam_user_policy_attachment" "test-attach" {
  user = aws_iam_user.lb.name
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_iam_group" "developers" {
  name = "aws_iam_group${var.test_prefix}${var.test_suffix}"
  path = "/users/"
}

resource "aws_iam_group_membership" "team" {
  name = "membership${var.test_prefix}${var.test_suffix}"

  users = [
    aws_iam_user.lb.name,
  ]

  group = aws_iam_group.developers.name
}


