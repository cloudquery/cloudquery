# Table: aws_ec2_vpn_gateways

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
|amazon_side_asn|Int|
|availability_zone|String|
|state|String|
|tags|JSON|
|type|String|
|vpc_attachments|JSON|
|vpn_gateway_id|String|