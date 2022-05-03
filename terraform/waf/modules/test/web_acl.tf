#

resource "aws_waf_ipset" "ipset" {
  name = "${var.prefix}wafipset"

  ip_set_descriptors {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_waf_rule" "wafrule" {
  name        = "${var.prefix}wafrule"
  metric_name = "${var.prefix}wafrule"

  predicates {
    data_id = aws_waf_ipset.ipset.id
    negated = false
    type    = "IPMatch"
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}wafrule"
    }
  )

  depends_on  = [aws_waf_ipset.ipset]
}

resource "aws_waf_web_acl" "waf_acl" {
  name        = "${var.prefix}wafwebacl"
  metric_name = "${var.prefix}wafwebacl"

  default_action {
    type = "ALLOW"
  }

  rules {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_waf_rule.wafrule.id
    type     = "REGULAR"
  }

  logging_configuration {
    log_destination = aws_kinesis_firehose_delivery_stream.extended_s3_stream.arn
    redacted_fields {
      field_to_match {
        type = "URI"
      }
    }
  }

  tags = merge(
    var.tags,
    {
      Name = "${var.prefix}wafwebacl"
    }
  )

  depends_on = [
    aws_waf_ipset.ipset,
    aws_waf_rule.wafrule,
  ]
}
