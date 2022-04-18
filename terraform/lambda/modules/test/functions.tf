# TF Module https://github.com/terraform-aws-modules/terraform-aws-lambda/blob/master/examples/complete/main.tf

module "lambda_function" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "~> 2.36"

  function_name = "${var.prefix}-lambda"
  description   = "${var.prefix}-lambda"
  handler       = "index.lambda_handler"
  runtime       = "python3.8"
  architectures = ["x86_64"]
  publish       = true

  source_path = "${path.root}/../modules/test/fixtures/python3.8-app1"

  store_on_s3 = true
  s3_bucket   = module.s3_bucket.s3_bucket_id
  s3_prefix   = "lambda-builds/"

  artifacts_dir = ".terraform/lambda-builds/"

  layers = [
    module.lambda_layer_local.lambda_layer_arn,
  ]

  environment_variables = {
    Hello      = "World"
    Serverless = "Terraform"
  }

  role_path   = "/tf-managed/"
  policy_path = "/tf-managed/"

  attach_dead_letter_policy = true
  dead_letter_target_arn    = aws_sqs_queue.dlq.arn

}




##################
# Extra resources
##################

module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "~> 3.0.1"

  bucket        = "cq-provider-${var.prefix}-lambda"
  force_destroy = true
}
