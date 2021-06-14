resource "aws_dx_gateway" "cq-provider-aws-dx-gateway" {
  name            = "${var.test_prefix}-${var.test_suffix}"
  amazon_side_asn = "64512"
}