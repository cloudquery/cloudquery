# Table: aws_ec2_vpn_connections

This table shows data for Amazon Elastic Compute Cloud (EC2) VPN Connections.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpnConnections.html

The composite primary key for this table is (**account_id**, **region**, **vpn_connection_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|category|`utf8`|
|core_network_arn|`utf8`|
|core_network_attachment_arn|`utf8`|
|customer_gateway_configuration|`utf8`|
|customer_gateway_id|`utf8`|
|gateway_association_state|`utf8`|
|options|`json`|
|routes|`json`|
|state|`utf8`|
|transit_gateway_id|`utf8`|
|type|`utf8`|
|vgw_telemetry|`json`|
|vpn_connection_id (PK)|`utf8`|
|vpn_gateway_id|`utf8`|