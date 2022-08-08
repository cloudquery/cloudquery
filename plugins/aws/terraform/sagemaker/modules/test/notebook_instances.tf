resource "aws_sagemaker_code_repository" "example" {
  code_repository_name = "${var.prefix}-sagemaker-cq-provider"

  git_config {
    repository_url = "https://github.com/cloudquery/cq-provider-aws.git"
  }
}

resource "aws_iam_role" "example" {
  name = "${var.prefix}-sagemaker-cq-provider"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "",
            "Effect": "Allow",
            "Principal": {
                "Service": "sagemaker.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"]
  tags = var.tags
}

resource "aws_sagemaker_notebook_instance" "example" {
  name          = "${var.prefix}-sagemaker-cq-provider"
  instance_type = "ml.t2.medium"
  role_arn      = aws_iam_role.example.arn

  tags = var.tags
}
