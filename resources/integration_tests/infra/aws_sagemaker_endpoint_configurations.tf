resource "aws_sagemaker_endpoint_configuration" "sagemaker_endpoint_configuration" {
  name = "sagemaker-endpoint-configuration-${var.test_prefix}${var.test_suffix}"

  production_variants {
    variant_name           = "variant-${var.test_prefix}${var.test_suffix}"
    model_name             = aws_sagemaker_model.sagemaker_model.name
    initial_instance_count = 1
    instance_type          = "ml.t2.medium"
  }

  tags = {
    Name = "sagemaker-endpoint-configuration-${var.test_prefix}${var.test_suffix}"
  }
}