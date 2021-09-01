resource "aws_apigatewayv2_api" "v2_api_1" {
  name                       = "apigwv2-api-${var.test_prefix}${var.test_suffix}"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_integration" "v2_integration_1" {
  api_id           = aws_apigatewayv2_api.v2_api_1.id
  integration_type = "HTTP_PROXY"

  integration_method = "ANY"
  integration_uri    = "https://example.com/{proxy}"
}

resource "aws_apigatewayv2_model" "v2_model_1" {
  api_id       = aws_apigatewayv2_api.v2_api_1.id
  content_type = "application/json"
  name         = "apigwv2model${var.test_suffix}"

  schema = <<EOF
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "ExampleModel",
  "type": "object",
  "properties": {
    "id": { "type": "string" }
  }
}
EOF
}

resource "aws_apigatewayv2_stage" "v2_stage_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  name   = "apigwv2-stage-${var.test_prefix}${var.test_suffix}"
}

resource "aws_apigatewayv2_route" "v2_route_1" {
  api_id    = aws_apigatewayv2_api.v2_api_1.id
  route_key = "GET /example/v1/test"

  target = "integrations/${aws_apigatewayv2_integration.v2_integration_1.id}"
}

resource "aws_apigatewayv2_route_response" "v2_route_response" {
  api_id             = aws_apigatewayv2_api.v2_api_1.id
  route_id           = aws_apigatewayv2_route.v2_route_1.id
  route_response_key = "$default"
}

resource "aws_apigatewayv2_deployment" "v2_deployment_1" {
  api_id      = aws_apigatewayv2_route.v2_route_1.api_id
  description = "apigwv2-dep-${var.test_prefix}${var.test_suffix}"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_apigatewayv2_authorizer" "v2_authorizer_1" {
  api_id          = aws_apigatewayv2_api.v2_api_1.id
  authorizer_type = "REQUEST"
  authorizer_uri  = aws_lambda_function.lambda_func.invoke_arn
  identity_sources = [
  "route.request.header.Auth"]
  name = "apigwv2-authorizer-${var.test_prefix}${var.test_suffix}"
}

resource "aws_iam_role" "v2_iam_role_1" {
  name = "apigwv2-role-${var.test_prefix}${var.test_suffix}"
  path = "/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "v2_role_policy" {
  name = "apigwv2-policy-${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.v2_iam_role_1.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "lambda:InvokeFunction",
      "Effect": "Allow",
      "Resource": "${aws_lambda_function.lambda_func.arn}"
    }
  ]
}
EOF
}

resource "aws_iam_role" "v2_iam_role_2" {
  name = "apigwv2-assume-role${var.test_prefix}${var.test_suffix}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_apigatewayv2_integration_response" "v2_response_1" {
  api_id                   = aws_apigatewayv2_api.v2_api_1.id
  integration_id           = aws_apigatewayv2_integration.v2_integration_1.id
  integration_response_key = "/200/"
}
