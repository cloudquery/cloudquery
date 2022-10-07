# Table: aws_ec2_customer_gateways



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
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