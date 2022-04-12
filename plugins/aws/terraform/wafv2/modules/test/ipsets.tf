resource "aws_wafv2_ip_set" "ipset1" {
  name               = "${var.prefix}-ipset"
  description        = "Example IP set"
  scope              = "REGIONAL"
  ip_address_version = "IPV4"
  addresses          = ["1.2.3.4/32", "5.6.7.8/32"]

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}

resource "aws_wafv2_ip_set" "ipset2" {
  name               = "${var.prefix}-ipset"
  description        = "Example IP set"
  scope              = "CLOUDFRONT"
  ip_address_version = "IPV4"
  addresses          = ["1.2.3.4/32", "5.6.7.8/32"]

  tags = {
    Tag1 = "Value1"
    Tag2 = "Value2"
  }
}
