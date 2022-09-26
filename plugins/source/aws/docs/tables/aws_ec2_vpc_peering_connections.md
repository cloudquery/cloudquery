# Table: aws_ec2_vpc_peering_connections


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|accepter_vpc_info|JSON|
|expiration_time|Timestamp|
|requester_vpc_info|JSON|
|status|JSON|
|tags|JSON|
|vpc_peering_connection_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|