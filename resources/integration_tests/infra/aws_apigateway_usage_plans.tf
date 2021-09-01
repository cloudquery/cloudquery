resource "aws_api_gateway_usage_plan" "api_gateway_usage_plan_1" {
  name         = "apigw-up-${var.test_prefix}${var.test_suffix}"
  description  = "my description"
  product_code = "MYCODE"

  api_stages {
    api_id = aws_api_gateway_rest_api.rest_api_example_1.id
    stage  = aws_api_gateway_stage.stage_1.stage_name
  }

  api_stages {
    api_id = aws_api_gateway_rest_api.rest_api_example_1.id
    stage  = aws_api_gateway_stage.stage_2.stage_name
  }

  quota_settings {
    limit  = 20
    offset = 2
    period = "WEEK"
  }

  throttle_settings {
    burst_limit = 5
    rate_limit  = 10
  }
}

resource "aws_api_gateway_usage_plan_key" "apigateway_usage_plan_key_1" {
  key_id        = aws_api_gateway_api_key.aws_apigateway_api_keys_key.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.api_gateway_usage_plan_1.id
}
