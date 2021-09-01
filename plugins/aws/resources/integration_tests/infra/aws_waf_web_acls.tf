resource "aws_waf_web_acl" "waf_web_acl_1" {
  depends_on = [
    aws_waf_ipset.waf_ipset_1,
    aws_waf_rule.waf_rule_1,
  ]
  name        = "waf-web-acl-${var.test_prefix}${var.test_suffix}"
  metric_name = "wafwebacl1"

  default_action {
    type = "ALLOW"
  }

  rules {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_waf_rule.waf_rule_1.id
    type     = "REGULAR"
  }
}
