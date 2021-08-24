resource "aws_route53_delegation_set" "route53_delegation_set" {
  reference_name = "route53_delegation_set${var.test_prefix}${var.test_suffix}"
}
