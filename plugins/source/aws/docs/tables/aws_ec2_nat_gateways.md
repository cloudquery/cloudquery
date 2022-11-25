# Table: aws_ec2_nat_gateways

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NatGateway.html

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
|connectivity_type|String|
|create_time|Timestamp|
|delete_time|Timestamp|
|failure_code|String|
|failure_message|String|
|nat_gateway_addresses|JSON|
|nat_gateway_id|String|
|provisioned_bandwidth|JSON|
|state|String|
|subnet_id|String|
|tags|JSON|
|vpc_id|String|