resource "aws_apigatewayv2_domain_name" "adnv2" {
  domain_name = "${var.test_prefix}${var.test_suffix}.com"

  domain_name_configuration {
    certificate_arn = aws_acm_certificate.adnv2.arn
    endpoint_type = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_apigatewayv2_api_mapping" "adnv2" {
  api_id = aws_apigatewayv2_api.adnv2.id
  domain_name = aws_apigatewayv2_domain_name.adnv2.id
  stage = aws_apigatewayv2_stage.adnv2.id
}

resource "aws_apigatewayv2_api" "adnv2" {
  name = "v2dn${var.test_prefix}${var.test_suffix}"
  protocol_type = "HTTP"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_stage" "adnv2" {
  api_id = aws_apigatewayv2_api.adnv2.id
  name = "v2stage${var.test_prefix}${var.test_suffix}"
}


resource "aws_route53_zone" "adnv2" {
  name = "${var.test_prefix}${var.test_suffix}.com"
}

resource "aws_acm_certificate" "adnv2" {
  domain_name = "${var.test_prefix}${var.test_suffix}.com"
  validation_method = "DNS"
}

resource "aws_acm_certificate_validation" "adnv2" {
  certificate_arn = aws_acm_certificate.adnv2.arn
  validation_record_fqdns = [for record in aws_route53_record.adnv2 : record.fqdn]
}

resource "aws_route53_record" "adnv2" {
  for_each = {
  for dvo in aws_acm_certificate.adnv2.domain_validation_options : dvo.domain_name => {
    name   = dvo.resource_record_name
    record = dvo.resource_record_value
    type   = dvo.resource_record_type
  }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id = aws_route53_zone.adnv2.zone_id

  depends_on = [aws_acm_certificate.adnv2]
}
