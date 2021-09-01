resource "aws_dx_lag" "aws_directconnect_virtual_interfaces_lag" {
  name                  = "dx-lag"
  connections_bandwidth = "1Gbps"
  location              = "EqDC2"
  force_destroy         = true
}


resource "time_sleep" "aws_directconnect_virtual_interfaces_wait_for_id" {
  depends_on = [
  aws_dx_lag.aws_directconnect_virtual_interfaces_lag]

  create_duration = "5m"

  triggers = {
    id = aws_dx_lag.aws_directconnect_virtual_interfaces_lag.id
  }
}

resource "aws_dx_public_virtual_interface" "aws_directconnect_virtual_interfaces_interface" {
  connection_id = time_sleep.aws_directconnect_virtual_interfaces_wait_for_id.triggers["id"]

  name           = "fx-pvif-${var.test_prefix}-${var.test_suffix}"
  vlan           = 4094
  address_family = "ipv4"
  bgp_asn        = 65352

  customer_address = "175.45.176.1/30"
  amazon_address   = "175.45.176.2/30"

  route_filter_prefixes = [
    "210.52.109.0/24",
    "175.45.176.0/22",
  ]

  depends_on = [
    time_sleep.aws_directconnect_virtual_interfaces_wait_for_id
  ]
}