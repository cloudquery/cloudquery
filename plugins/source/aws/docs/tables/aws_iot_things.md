# Table: aws_iot_things



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
|principals|StringArray|
|arn (PK)|String|
|attributes|JSON|
|thing_name|String|
|thing_type_name|String|
|version|Int|