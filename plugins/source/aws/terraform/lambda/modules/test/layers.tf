#################################
# Lambda Layer (storing locally)
#################################

module "lambda_layer_local" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "~> 2.36"

  create_layer = true

  layer_name               = "${var.prefix}-lambda"
  description              = "${var.prefix}-lambda lambda layer (deployed from local)"
  compatible_runtimes      = ["python3.8"]
  compatible_architectures = ["arm64"]

  source_path = "${path.root}/../modules/test/fixtures/python3.8-app1"
}

##################
# Extra resources
##################

resource "aws_sqs_queue" "dlq" {
  name = "${var.prefix}-lambda"
}

