resource "aws_api_gateway_rest_api" "rest_api_example_1" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "apigwv1-api-${var.test_prefix}${var.test_suffix}"

  endpoint_configuration {
    types = [
      "REGIONAL"
    ]
  }
}

resource "aws_api_gateway_deployment" "deployment_example_1" {
  rest_api_id = aws_api_gateway_rest_api.rest_api_example_1.id
  variables = {
    test : "test"
  }
  description = "apigwv1-dep-${var.test_prefix}${var.test_suffix}"

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.rest_api_example_1.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "stage_1" {
  deployment_id = aws_api_gateway_deployment.deployment_example_1.id
  rest_api_id   = aws_api_gateway_rest_api.rest_api_example_1.id
  stage_name    = "apigwv1-stage-${var.test_prefix}${var.test_suffix}"
  tags = {
    "hello" = "world"
  }
}

resource "aws_api_gateway_stage" "stage_2" {
  deployment_id = aws_api_gateway_deployment.deployment_example_1.id
  rest_api_id   = aws_api_gateway_rest_api.rest_api_example_1.id
  stage_name    = "apigwv1-stage2-${var.test_prefix}${var.test_suffix}"
  tags = {
    "hello" = "world1"
  }
}

resource "aws_api_gateway_resource" "gateway_resource_1" {
  rest_api_id = aws_api_gateway_rest_api.rest_api_example_1.id
  parent_id   = aws_api_gateway_rest_api.rest_api_example_1.root_resource_id
  path_part   = "gateway_resource_1"
}

resource "aws_api_gateway_method" "gateway_method_1" {
  rest_api_id   = aws_api_gateway_rest_api.rest_api_example_1.id
  resource_id   = aws_api_gateway_resource.gateway_resource_1.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "gateway_integration_1" {
  rest_api_id = aws_api_gateway_rest_api.rest_api_example_1.id
  resource_id = aws_api_gateway_resource.gateway_resource_1.id
  http_method = aws_api_gateway_method.gateway_method_1.http_method
  type        = "MOCK"
  // cache_key_parameters = ["method.request.path.param"]
  cache_namespace      = "foobar"
  timeout_milliseconds = 29000

  request_parameters = {
    "integration.request.header.X-Authorization" = "'static'"
  }

  # Transforms the incoming XML request to JSON
  request_templates = {
    "application/xml" = <<EOF
{
   "body" : $input.json('$')
}
EOF
  }
}

resource "aws_api_gateway_authorizer" "gateway_authorizer_1" {
  name                             = "apigwv1-authorizer-${var.test_prefix}${var.test_suffix}"
  rest_api_id                      = aws_api_gateway_rest_api.rest_api_example_1.id
  authorizer_uri                   = aws_lambda_function.lambda_func.invoke_arn
  authorizer_credentials           = aws_iam_role.invocation_role.arn
  authorizer_result_ttl_in_seconds = 500
  type                             = "TOKEN"
}

resource "aws_iam_role" "invocation_role" {
  name = "apigwv1-api-invocation-role-${var.test_prefix}${var.test_suffix}"
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

resource "aws_iam_role_policy" "invocation_policy" {
  name = "apigwv1-api-invocation-policy-${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.invocation_role.id

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

resource "aws_iam_role" "lambda" {
  name = "apigwv1-api-assume-role-${var.test_prefix}${var.test_suffix}"

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

resource "aws_api_gateway_model" "gateway_model_1" {
  rest_api_id  = aws_api_gateway_rest_api.rest_api_example_1.id
  name         = "apigwv1apimodel${var.test_suffix}"
  description  = "a JSON schema"
  content_type = "application/json"

  schema = <<EOF
{
  "type": "object"
}
EOF
}

resource "aws_api_gateway_request_validator" "request_validator_1" {
  name                        = "apigwv1-req-validation-${var.test_prefix}${var.test_suffix}"
  rest_api_id                 = aws_api_gateway_rest_api.rest_api_example_1.id
  validate_request_body       = true
  validate_request_parameters = true
}

resource "aws_api_gateway_documentation_part" "documentation_part_1" {
  location {
    type   = "METHOD"
    method = "GET"
    path   = "/example"
  }

  properties  = "{\"description\":\"Example description\"}"
  rest_api_id = aws_api_gateway_rest_api.rest_api_example_1.id
}

resource "aws_api_gateway_documentation_version" "documentation_version_1" {
  version     = "example_version"
  rest_api_id = aws_api_gateway_rest_api.rest_api_example_1.id
  description = "Example description"
  depends_on = [
  aws_api_gateway_documentation_part.documentation_part_1]
}
