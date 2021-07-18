
# Table: gcp_compute_vpn_gateway_vpn_interfaces
A VPN gateway interface
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpn_gateway_cq_id|uuid|Unique ID of gcp_compute_vpn_gateways table (FK)|
|vpn_gateway_id|text||
|id|text|The numeric ID of this VPN gateway interface|
|interconnect_attachment|text|URL of the interconnect attachment resource When the value of this field is present, the VPN Gateway will be used for IPsec-encrypted Cloud Interconnect; all Egress or Ingress traffic for this VPN Gateway interface will go through the specified interconnect attachment resource Not currently available in all Interconnect locations|
|ip_address|text|The external IP address for this VPN gateway interface|
