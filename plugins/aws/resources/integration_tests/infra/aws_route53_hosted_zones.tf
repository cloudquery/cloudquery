resource "aws_route53_zone" "route53_zone_dev" {
  name = "dev.${var.test_prefix}${var.test_suffix}.com"
  tags = {
    Environment = "dev"
  }
}

resource "aws_route53_record" "route53_zone_record" {
  zone_id = aws_route53_zone.route53_zone_dev.zone_id
  name    = "dev-1.${var.test_prefix}${var.test_suffix}.com"
  type    = "NS"
  ttl     = "30"
  records = aws_route53_zone.route53_zone_dev.name_servers
}
