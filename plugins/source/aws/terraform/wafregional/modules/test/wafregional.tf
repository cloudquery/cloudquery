resource "aws_wafregional_ipset" "ipset1" {
  name = "${var.prefix}_ipset1"

  ip_set_descriptor {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_wafregional_rate_based_rule" "rate_rule1" {
  depends_on  = [aws_wafregional_ipset.ipset1]
  name        = "${var.prefix}_rate_rule1"
  metric_name = "${var.prefix}ratemetric1"

  rate_key   = "IP"
  rate_limit = 100

  predicate {
    data_id = aws_wafregional_ipset.ipset1.id
    negated = false
    type    = "IPMatch"
  }

  tags = {
    "key" = "rate based rule"
  }
}

resource "aws_wafregional_rule" "rule1" {
  name        = "${var.prefix}_rule1"
  metric_name = "${var.prefix}metric1"

  predicate {
    type    = "IPMatch"
    data_id = aws_wafregional_ipset.ipset1.id
    negated = false
  }

  tags = {
    "key" = "rule"
  }
}

resource "aws_wafregional_rule_group" "rule_group" {
  name        = "${var.prefix}_rule_group"
  metric_name = "${var.prefix}groupmetric"

  activated_rule {
    action {
      type = "COUNT"
    }

    priority = 50
    rule_id  = aws_wafregional_rule.rule1.id
  }

  tags = {
    "key" = "rule group"
  }
}

resource "aws_wafregional_web_acl" "wafacl" {
  name        = "${var.prefix}_web_acl"
  metric_name = "${var.prefix}webaclmetric"

  default_action {
    type = "ALLOW"
  }

  rule {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_wafregional_rule.rule1.id
    type     = "REGULAR"
  }

  tags = {
    "key" = "web acl"
  }
}
