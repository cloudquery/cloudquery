resource "aws_wafv2_rule_group" "wafv2_rule_group" {
  name     = "${var.prefix}-rule-group"
  scope    = "REGIONAL"
  capacity = 2

  rule {
    name     = "${var.prefix}-rule-1"
    priority = 1

    action {
      block {
        custom_response {
          custom_response_body_key  = "blocked_request_custom_response"
          response_code             = 429
        }
      }
    }

    statement {

      geo_match_statement {
        country_codes = ["US", "NL"]
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = false
      metric_name                = "${var.prefix}-rule-metric-1"
      sampled_requests_enabled   = false
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "${var.prefix}-rule-group-metric-1"
    sampled_requests_enabled   = false
  }

  custom_response_body {
    key           = "blocked_request_custom_response"
    content       = "Rate Limit Exceeded"
    content_type  = "TEXT_PLAIN"
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-rule-group"
    }
  )
}