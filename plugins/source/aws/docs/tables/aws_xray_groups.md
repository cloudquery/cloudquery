# Table: aws_xray_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|filter_expression|String|
|group_name|String|
|insights_configuration|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|