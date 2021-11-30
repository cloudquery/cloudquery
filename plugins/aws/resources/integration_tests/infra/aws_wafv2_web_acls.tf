resource "aws_wafv2_web_acl" "wafv2_web_acl_1" {
  name        = "wafv2-web-acl-${var.test_prefix}${var.test_suffix}"
  description = "Example of a managed rule."
  scope       = "REGIONAL"

  default_action {
    allow {}
  }

  rule {
    name     = "rule-1"
    priority = 1

    override_action {
      count {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesCommonRuleSet"
        vendor_name = "AWS"

        excluded_rule {
          name = "SizeRestrictions_QUERYSTRING"
        }

        excluded_rule {
          name = "NoUserAgent_HEADER"
        }

        scope_down_statement {
          geo_match_statement {
            country_codes = ["US", "NL"]
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = false
      metric_name                = "friendly-rule-metric-name"
      sampled_requests_enabled   = false
    }
  }

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "friendly-metric-name"
    sampled_requests_enabled   = false
  }
}

resource "aws_wafv2_web_acl_logging_configuration" "wafv2_web_acl_logging_configuration" {
  log_destination_configs = [aws_kinesis_firehose_delivery_stream.kinesis_firehose_delivery_stream.arn]
  resource_arn            = aws_wafv2_web_acl.wafv2_web_acl_1.arn
  redacted_fields {
    single_header {
      name = "user-agent"
    }
  }

  depends_on = [aws_kinesis_firehose_delivery_stream.kinesis_firehose_delivery_stream]
}