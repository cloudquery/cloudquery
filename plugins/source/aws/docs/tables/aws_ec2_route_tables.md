# Table: aws_ec2_route_tables

This table shows data for Amazon Elastic Compute Cloud (EC2) Route Tables.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTable.html

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
|associations|`json`|
|owner_id|`utf8`|
|propagating_vgws|`json`|
|route_table_id|`utf8`|
|routes|`json`|
|vpc_id|`utf8`|