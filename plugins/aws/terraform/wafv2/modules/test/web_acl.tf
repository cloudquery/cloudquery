resource "aws_wafv2_web_acl" "wafv2_web_acl" {
  name        = "${var.prefix}-web-acl-managed-rule"
  description = "Example of a managed rule."
  scope       = "REGIONAL"

  default_action {
    allow {}
  }

  rule {
    name     = "${var.prefix}-rule-1"
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
      metric_name                = "${var.prefix}-web-acl-rule-metric-1"
      sampled_requests_enabled   = false
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "${var.prefix}-web-acl-metric-1"
    sampled_requests_enabled   = false
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}-web-acl-managed-rule"
    }
  )
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

  depends_on = [module.wafv2-web-acl-logging]
}