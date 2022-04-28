resource "aws_wafv2_rule_group" "wafv2_rule_group" {
  name     = "${var.prefix}-rule-group"
  scope    = "REGIONAL"
  capacity = 2

  rule {
    name     = "${var.prefix}-rule-1"
    priority = 1

    action {
      allow {}
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

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-rule-group"
    }
  )
}