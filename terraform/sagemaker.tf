resource "aws_sagemaker_notebook_instance" "sagemaker_notebook" {
  name          = "cq-provider-aws-sagemaker-notebook"
  instance_type = "ml.t2.medium"
  role_arn      = "arn:aws:iam::707066037989:role/cq-provier-aws-sagemaker-role"

  tags = {
    Name = "cq-provider-aws-sagemaker-notebook"
  }
}
