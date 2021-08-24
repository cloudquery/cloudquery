resource "aws_sns_topic" "lambda_func_sns_topic_user_updates" {
  name = "lambda_func_user-updates-topic${var.test_prefix}${var.test_suffix}"
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

resource "aws_sns_topic" "lambda_func_sns_topic_errors_topic" {
  name = "lambda_func_errors-topic${var.test_prefix}${var.test_suffix}"
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

resource "aws_lambda_permission" "lambda_func_permission_with_sns" {
  statement_id = "AllowExecutionFromSNS"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_func.function_name
  principal = "sns.amazonaws.com"
  source_arn = aws_sns_topic.lambda_func_sns_topic_user_updates.arn
}

resource "aws_lambda_alias" "lambda_func_alias" {
  name = "lambda_func_alias-${var.test_prefix}${var.test_suffix}"
  description = "a sample description"
  function_name = aws_lambda_function.lambda_func.function_name
  function_version = "$LATEST"
}

resource "aws_lambda_function_event_invoke_config" "lambda_func_invoke_config" {
  function_name = aws_lambda_alias.lambda_func_alias.function_name

  destination_config {
    on_failure {
      destination = aws_sns_topic.lambda_func_sns_topic_errors_topic.arn
    }

    on_success {
      destination = aws_sns_topic.lambda_func_sns_topic_errors_topic.arn
    }
  }
}

# See also the following AWS managed policy: AWSLambdaBasicExecutionRole
resource "aws_iam_policy" "lambda_func_iam_policy_publish" {
  name = "lambda_${var.test_prefix}${var.test_suffix}"
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

resource "aws_iam_role_policy_attachment" "lambda_func_policy_logs" {
  role = aws_iam_role.lambda_func_iam_role.name
  policy_arn = aws_iam_policy.lambda_func_iam_policy_publish.arn
}

resource "aws_iam_role" "lambda_func_iam_role" {
  name = "lambda_func_role_${var.test_prefix}${var.test_suffix}"

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

resource "aws_lambda_function" "lambda_func" {
  filename = data.archive_file.lambda_func_zip_inline.output_path
  source_code_hash = data.archive_file.lambda_func_zip_inline.output_base64sha256
  function_name = "function_${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.lambda_func_iam_role.arn
  handler = "exports.example"
  runtime = "nodejs12.x"
  publish = true

  environment {
    variables = {
      foo = "bar"
    }
  }
}

data "archive_file" "lambda_func_zip_inline" {
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
