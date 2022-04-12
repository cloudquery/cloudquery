resource "aws_wafv2_regex_pattern_set" "regex1" {
  name        = "${var.prefix}-regex-set1"
  description = "Example regex pattern set 1"
  scope       = "REGIONAL"

  regular_expression {
    regex_string = "one"
  }

  regular_expression {
    regex_string = "two"
  }

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}

resource "aws_wafv2_regex_pattern_set" "regex2" {
  name        = "${var.prefix}-regex-set2"
  description = "Example regex pattern set 2"
  scope       = "CLOUDFRONT"

  regular_expression {
    regex_string = "one"
  }

  regular_expression {
    regex_string = "two"
  }

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}
