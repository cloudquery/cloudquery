resource "aws_apigatewayv2_api" "example" {
  name = "v2${var.test_prefix}${var.test_suffix}"
  protocol_type = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_integration" "example" {
  api_id = aws_apigatewayv2_api.example.id
  integration_type = "HTTP_PROXY"

  integration_method = "ANY"
  integration_uri = "https://example.com/{proxy}"
}

resource "aws_apigatewayv2_model" "example" {
  api_id = aws_apigatewayv2_api.example.id
  content_type = "application/json"
  name = "v2model${var.test_prefix}${var.test_suffix}"

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

resource "aws_apigatewayv2_stage" "example" {
  api_id = aws_apigatewayv2_api.example.id
  name = "v2stage${var.test_prefix}${var.test_suffix}"
}

resource "aws_apigatewayv2_route" "example" {
  api_id = aws_apigatewayv2_api.example.id
  route_key = "GET /example/v1/test"

  target = "integrations/${aws_apigatewayv2_integration.example.id}"
}

resource "aws_apigatewayv2_route_response" "example" {
  api_id = aws_apigatewayv2_api.example.id
  route_id = aws_apigatewayv2_route.example.id
  route_response_key = "$default"
}

resource "aws_apigatewayv2_deployment" "example" {
  api_id = aws_apigatewayv2_route.example.api_id
  description = "Example deployment"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_apigatewayv2_authorizer" "example" {
  api_id = aws_apigatewayv2_api.example.id
  authorizer_type = "REQUEST"
  authorizer_uri = aws_lambda_function.authorizer_v2.invoke_arn
  identity_sources = [
    "route.request.header.Auth"]
  name = "example-authorizer"
}


resource "aws_iam_role" "invocation_role" {
  name = "apiv2${aws_apigatewayv2_integration.example.id}"
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
  name = "v2apipolicy${aws_apigatewayv2_integration.example.id}"
  role = aws_iam_role.invocation_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "lambda:InvokeFunction",
      "Effect": "Allow",
      "Resource": "${aws_lambda_function.authorizer_v2.arn}"
    }
  ]
}
EOF
}

resource "aws_iam_role" "lambda" {
  name = "v2api_lambda_role${aws_apigatewayv2_integration.example.id}"

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

resource "aws_lambda_function" "authorizer_v2" {
  filename = data.archive_file.lambda_zip_inline.output_path
  source_code_hash = data.archive_file.lambda_zip_inline.output_base64sha256
  function_name = "v2authorizer${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.lambda.arn
  handler = "exports.example"
  runtime = "nodejs12.x"

}

resource "aws_apigatewayv2_integration_response" "example" {
  api_id = aws_apigatewayv2_api.example.id
  integration_id = aws_apigatewayv2_integration.example.id
  integration_response_key = "/200/"
}


data "archive_file" "lambda_zip_inline" {
  type = "zip"
  output_path = "./tmp/lambda_zip_inline.zip"
  source {
    content = <<EOF
module.exports.handler = async (event, context, callback) => {
	const what = "world";
	const response = `Hello $${what}!`;
	callback(null, response);
};
EOF
    filename = "main.js"
  }
}
