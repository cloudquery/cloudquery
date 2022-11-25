# Table: aws_ec2_route_tables

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTable.html

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
|associations|JSON|
|owner_id|String|
|propagating_vgws|JSON|
|route_table_id|String|
|routes|JSON|
|tags|JSON|
|vpc_id|String|