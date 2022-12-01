# Table: aws_ec2_customer_gateways

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CustomerGateway.html

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
|bgp_asn|String|
|certificate_arn|String|
|customer_gateway_id|String|
|device_name|String|
|ip_address|String|
|state|String|
|tags|JSON|
|type|String|