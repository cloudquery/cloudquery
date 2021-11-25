resource "aws_iam_role" "sagemaker_notebook_instance_iam_role" {
  name = "sagemaker-notebook-instance-iam-role-${var.test_prefix}${var.test_suffix}"
  path = "/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "sagemaker.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_sagemaker_notebook_instance" "sagemaker_notebook_instance" {
  name          = "sagemaker-${var.test_prefix}${var.test_suffix}"
  role_arn      = aws_iam_role.sagemaker_notebook_instance_iam_role.arn
  instance_type = "ml.t2.medium"

  tags = {
    Name = "sagemaker-${var.test_prefix}${var.test_suffix}"
  }
}