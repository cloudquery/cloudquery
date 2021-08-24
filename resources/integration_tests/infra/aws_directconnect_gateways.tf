resource "aws_dx_gateway" "integ-aws-dx-gateway" {
  name            = "dx-gateway${var.test_prefix}-${var.test_suffix}"
  amazon_side_asn = "64512"
}