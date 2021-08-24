resource "aws_cloudwatch_log_metric_filter" "aws_cloudwatch_log_metric_filter" {
  name           = "aws_cloudwatch_log_metric_filter_${var.test_prefix}${var.test_suffix}"
  pattern        = ""
  log_group_name = aws_cloudwatch_log_group.aws_cloudwatch_log_metric_filter_group.name

  metric_transformation {
    name      = "aws_cloudwatch_log_metric_filter_name"
    namespace = "YourNamespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_log_group" "aws_cloudwatch_log_metric_filter_group" {
  name = "MyApp${var.test_prefix}${var.test_suffix}/access.log"
}