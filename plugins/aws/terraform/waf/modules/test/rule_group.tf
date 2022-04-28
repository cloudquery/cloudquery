resource "aws_waf_rule" "wafrule_2" {
  name        = "${var.prefix}wafrule2"
  metric_name = "${var.prefix}wafrule2"

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}wafrule2"
    }
  )
}

resource "aws_waf_rule_group" "waf_rule_group" {
  name        = "${var.prefix}wafrulegroup"
  metric_name = "${var.prefix}wafrulegroup"

  activated_rule {
    action {
      type = "COUNT"
    }

    priority = 50
    rule_id  = aws_waf_rule.wafrule_2.id
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}wafrulegroup"
    }
  )
}