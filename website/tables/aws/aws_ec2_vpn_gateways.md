# Table: aws_ec2_vpn_gateways

This table shows data for Amazon Elastic Compute Cloud (EC2) VPN Gateways.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnGateway.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|amazon_side_asn|Int|
|availability_zone|String|
|state|String|
|type|String|
|vpc_attachments|JSON|
|vpn_gateway_id|String|