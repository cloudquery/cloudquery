# Table: aws_ec2_vpn_gateways


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|