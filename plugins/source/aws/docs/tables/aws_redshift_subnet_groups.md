# Table: aws_redshift_subnet_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|cluster_subnet_group_name|String|
|description|String|
|subnet_group_status|String|
|subnets|JSON|
|vpc_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|