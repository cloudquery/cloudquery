
resource "aws_route53_zone" "hz" {
  name              = "${var.test_prefix}${var.test_suffix}.com"
  delegation_set_id = aws_route53_delegation_set.main.id
  tags = {
    Environment = "prod"
    Contributor = "andrii"
  }
}


resource "aws_route53_zone" "hzdev" {
  name              = "dev.${var.test_prefix}${var.test_suffix}.com"
  tags = {
    Environment = "dev"
  }
}

resource "aws_route53_record" "hz" {
  zone_id = aws_route53_zone.hz.zone_id
  name    = "dev.${var.test_prefix}${var.test_suffix}.com"
  type    = "NS"
  ttl     = "30"
  records = aws_route53_zone.hzdev.name_servers
}
