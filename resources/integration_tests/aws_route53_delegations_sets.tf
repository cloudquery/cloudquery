
resource "aws_route53_delegation_set" "main" {
  reference_name = "aws_route53_delegation_set${var.test_prefix}${var.test_suffix}"
}
