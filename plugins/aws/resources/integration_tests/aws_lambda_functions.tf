resource "aws_sns_topic" "user_updates" {
  name = "user-updates-topic"
  delivery_policy = <<EOF
{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
  }
}
EOF
}

resource "aws_sns_topic" "errors_topic" {
  name = "errors-topic"
  delivery_policy = <<EOF
{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
  }
}
EOF
}


resource "aws_lambda_permission" "with_sns" {
  statement_id = "AllowExecutionFromSNS"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.test_lambda.function_name
  principal = "sns.amazonaws.com"
  source_arn = aws_sns_topic.user_updates.arn
}


resource "aws_lambda_alias" "test_alias" {
  name = "testalias"
  description = "a sample description"
  function_name = aws_lambda_function.test_lambda.function_name
  function_version = "$LATEST"
}

//resource "aws_lambda_provisioned_concurrency_config" "example" {
//  function_name = aws_lambda_function.test_lambda.function_name
//  provisioned_concurrent_executions = 1
//  qualifier = aws_lambda_function.test_lambda.version
//  timeouts {
//    create = "1m"
//  }
//}

resource "aws_lambda_function_event_invoke_config" "example" {
  function_name = aws_lambda_alias.test_alias.function_name

  destination_config {
    on_failure {
      destination = aws_sns_topic.errors_topic.arn
    }

    on_success {
      destination = aws_sns_topic.errors_topic.arn
    }
  }
}


# See also the following AWS managed policy: AWSLambdaBasicExecutionRole
resource "aws_iam_policy" "lambda_publish" {
  name = "lambda_logging"
  path = "/"
  description = "IAM policy for publishing from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sns:Publish",
      "Resource": "arn:aws:sns:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.lambda_publish.arn
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}

//dont know how to add layer version policy
resource "aws_lambda_layer_version" "lambda_layer" {
  filename = data.archive_file.lambda_zip_inline.output_path
  layer_name = "test_lambda_layer"

  compatible_runtimes = [
    "nodejs12.x"]
}


resource "aws_lambda_function" "test_lambda" {
  filename = data.archive_file.lambda_zip_inline.output_path
  source_code_hash = data.archive_file.lambda_zip_inline.output_base64sha256
  function_name = "test_function_${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.iam_for_lambda.arn
  handler = "exports.example"
  runtime = "nodejs12.x"
  publish = true

  environment {
    variables = {
      foo = "bar"
    }
  }

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