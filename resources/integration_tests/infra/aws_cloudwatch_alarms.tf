resource "aws_cloudwatch_metric_alarm" "cloudwatch_alarm" {
  alarm_name                = "cw-alarm${var.test_prefix}-${var.test_suffix}"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = "2"
  threshold                 = "10"
  alarm_description         = "Request error rate has exceeded 10%"
  insufficient_data_actions = []

  metric_query {
    id          = "e1"
    expression  = "m2/m1*100"
    label       = "Error Rate"
    return_data = "true"
  }

  metric_query {
    id = "m1"

    metric {
      metric_name = "RequestCount"
      namespace   = "AWS/ApplicationELB"
      period      = "120"
      stat        = "Sum"
      unit        = "Count"

      dimensions = {
        LoadBalancer = "app/web"
      }
    }
  }

  metric_query {
    id = "m2"

    metric {
      metric_name = "HTTPCode_ELB_5XX_Count"
      namespace   = "AWS/ApplicationELB"
      period      = "120"
      stat        = "Sum"
      unit        = "Count"

      dimensions = {
        LoadBalancer = "app/web"
      }
    }
  }
}

resource "aws_cloudwatch_log_metric_filter" "cloudwatch_log_filter" {
  name           = "cw-log-filter${var.test_prefix}-${var.test_suffix}"
  pattern        = ""
  log_group_name = aws_cloudwatch_log_group.cloudwatch_log_group.name

  metric_transformation {
    name      = "EventCount"
    namespace = "YourNamespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_log_group" "cloudwatch_log_group" {
  name = "cw-log-g${var.test_prefix}-${var.test_suffix}"
}