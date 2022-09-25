# Table: aws_ec2_route_tables


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|