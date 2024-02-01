# Table: aws_ec2_vpc_peering_connections

This table shows data for Amazon Elastic Compute Cloud (EC2) VPC Peering Connections.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnection.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|accepter_vpc_info|`json`|
|expiration_time|`timestamp[us, tz=UTC]`|
|requester_vpc_info|`json`|
|status|`json`|
|vpc_peering_connection_id|`utf8`|