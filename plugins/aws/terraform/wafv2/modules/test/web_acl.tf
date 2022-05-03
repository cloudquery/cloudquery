#

resource "aws_wafv2_regex_pattern_set" "example" {
  name        = "example"
  description = "Example regex pattern set"
  scope       = "REGIONAL"

  regular_expression {
    regex_string = "one"
  }
}

resource "aws_wafv2_web_acl" "wafv2_web_acl" {
  name        = "${var.prefix}-web-acl-managed-rule"
  description = "Example of a managed rule."
  scope       = "REGIONAL"

  default_action {
    allow {}
  }

  custom_response_body {
    key          = "blocked_request_custom_response"
    content      = "Rate Limit Exceeded"
    content_type = "TEXT_PLAIN"
  }

  rule {
    name     = "${var.prefix}-rule-1"
    priority = 2

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
      metric_name                = "${var.prefix}-web-acl-rule-metric-1"
      sampled_requests_enabled   = false
    }
  }

  rule {
    name     = "rate_limit_mt_external_webhook"
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
      rate_based_statement {
        limit              = 100   # 100 per span of 5 mins - minimum value
        aggregate_key_type = "IP"

        scope_down_statement {
          regex_pattern_set_reference_statement {
            arn = aws_wafv2_regex_pattern_set.example.arn
            text_transformation {
              priority = 1
              type     = "NONE"
            }
            field_to_match {
              uri_path {}
            }
          }
        }
      }

    }
    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "wafv2_alb_mt_external_webhook_acl_rule_vis"
      sampled_requests_enabled   = false
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = true
    metric_name                = "wafv2_alb_mt_external_webhook_acl_vis"
    sampled_requests_enabled   = false
  }
}

module "wafv2-web-acl-logging" {
  source        = "terraform-aws-modules/s3-bucket/aws"
  version       = "3.0.1"
  bucket        = "aws-waf-logs-${var.prefix}-s3-wafv2-web-acl-logging"
  force_destroy = true
}

resource "aws_wafv2_web_acl_logging_configuration" "web_acl_logging_configuration" {
  log_destination_configs = [module.wafv2-web-acl-logging.s3_bucket_arn]
  resource_arn            = aws_wafv2_web_acl.wafv2_web_acl.arn
  redacted_fields {
    single_header {
      name = "user-agent"
    }
  }

  logging_filter {
    default_behavior = "KEEP"

    filter {
      behavior = "DROP"

      condition {
        action_condition {
          action = "COUNT"
        }
      }

      requirement = "MEETS_ALL"
    }

    filter {
      behavior = "KEEP"

      condition {
        action_condition {
          action = "ALLOW"
        }
      }

      requirement = "MEETS_ANY"
    }
  }

  depends_on = [module.wafv2-web-acl-logging]
}