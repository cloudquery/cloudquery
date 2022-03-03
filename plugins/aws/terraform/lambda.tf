# TF Module https://github.com/terraform-aws-modules/terraform-aws-lambda/blob/master/examples/complete/main.tf

module "lambda_function" {
  source = "terraform-aws-modules/lambda/aws"
  version = "~> 2.34"

  function_name = "cq-provider-aws-lambda"
  description   = "cq-provider-aws-lambda"
  handler       = "index.lambda_handler"
  runtime       = "python3.8"
  architectures = ["x86_64"]
  publish       = true

  source_path = "${path.root}/fixtures/python3.8-app1"

  store_on_s3 = true
  s3_bucket   = module.cq_provider_aws_s3.s3_bucket_id
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


#################################
# Lambda Layer (storing locally)
#################################

module "lambda_layer_local" {
  source = "terraform-aws-modules/lambda/aws"
  version = "~> 2.34"

  create_layer = true

  layer_name               = "cq-provider-aws-layer-local"
  description              = "cq-provider-aws lambda layer (deployed from local)"
  compatible_runtimes      = ["python3.8"]
  compatible_architectures = ["arm64"]

  source_path = "${path.root}/fixtures/python3.8-app1"
}

##################
# Extra resources
##################

resource "random_pet" "this" {
  length = 2
}

resource "aws_sqs_queue" "dlq" {
  name = random_pet.this.id
}

