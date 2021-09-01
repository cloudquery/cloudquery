resource "aws_waf_rule_group" "waf_rule_group_1" {
  name        = "waf-rg-${var.test_prefix}${var.test_suffix}"
  metric_name = "wafrulegroup1"

  activated_rule {
    action {
      type = "COUNT"
    }

    priority = 50
    rule_id  = aws_waf_rule.waf_rule_1.id
  }
}
