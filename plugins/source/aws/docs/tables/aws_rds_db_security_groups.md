# Table: aws_rds_db_security_groups



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|db_security_group_description|String|
|db_security_group_name|String|
|e_c2_security_groups|JSON|
|ip_ranges|JSON|
|owner_id|String|
|vpc_id|String|