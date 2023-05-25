# Table: aws_ec2_route_tables

This table shows data for Amazon Elastic Compute Cloud (EC2) Route Tables.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTable.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|associations|json|
|owner_id|utf8|
|propagating_vgws|json|
|route_table_id|utf8|
|routes|json|
|vpc_id|utf8|