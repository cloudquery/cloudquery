resource "aws_iam_policy" "iam_policy_policy" {
  name = "iam_policy_${var.test_prefix}${var.test_suffix}"

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
