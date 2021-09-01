resource "aws_waf_ipset" "waf_ipset_1" {
  name = "waf-ipset-${var.test_prefix}${var.test_suffix}"

  ip_set_descriptors {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "waf_rule_1" {
  depends_on  = [aws_waf_ipset.waf_ipset_1]
  name        = "waf-rule-${var.test_prefix}${var.test_suffix}"
  metric_name = "wafrule1"

  predicates {
    data_id = aws_waf_ipset.waf_ipset_1.id
    negated = false
    type    = "IPMatch"
  }
}
