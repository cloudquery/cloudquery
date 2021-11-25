resource "aws_sagemaker_model" "sagemaker_model" {
  name               = "sagemaker-model-${var.test_prefix}${var.test_suffix}"
  execution_role_arn = aws_iam_role.sagemaker_model_iam_role.arn

  primary_container {
    image = data.aws_sagemaker_prebuilt_ecr_image.sagemaker_model_ecr_image.registry_path
  }

  tags = {
    Name = "sagemaker-model-${var.test_prefix}${var.test_suffix}"
  }
}

resource "aws_iam_role" "sagemaker_model_iam_role" {
  assume_role_policy = data.aws_iam_policy_document.sagemaker_model_assume_role.json
}

data "aws_iam_policy_document" "sagemaker_model_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

data "aws_sagemaker_prebuilt_ecr_image" "sagemaker_model_ecr_image" {
  repository_name = "kmeans"
}