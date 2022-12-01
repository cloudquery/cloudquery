# Table: aws_ec2_vpc_peering_connections

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnection.html

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
|accepter_vpc_info|JSON|
|expiration_time|Timestamp|
|requester_vpc_info|JSON|
|status|JSON|
|tags|JSON|
|vpc_peering_connection_id|String|