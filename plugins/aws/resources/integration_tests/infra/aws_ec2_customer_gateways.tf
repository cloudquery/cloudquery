resource "aws_customer_gateway" "aws_ec2_customer_gateways_gw" {
  bgp_asn    = 65000
  ip_address = "172.83.124.10"
  type       = "ipsec.1"

  tags = {
    Name = "ec2-cgw-${var.test_prefix}-${var.test_suffix}"
  }
}