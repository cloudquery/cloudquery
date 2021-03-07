resource "aws_cloudwatch_event_rule" "scan_schedule" {
  name = "Cloudquery us-east-1 scan"
  description = "Run cloudquery everyday on us-east-1 resources"

  schedule_expression = "rate(1 day)"
}

resource "aws_cloudwatch_event_target" "sns" {
  rule      = aws_cloudwatch_event_rule.scan_schedule.name
  arn       = aws_lambda_function.cloudquery.arn
  input     = file("tasks/us-east-1/input.json")
}