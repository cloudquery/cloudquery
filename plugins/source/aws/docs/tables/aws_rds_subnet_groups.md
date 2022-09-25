# Table: aws_rds_subnet_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|db_subnet_group_description|String|
|db_subnet_group_name|String|
|subnet_group_status|String|
|subnets|JSON|
|supported_network_types|StringArray|
|vpc_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|