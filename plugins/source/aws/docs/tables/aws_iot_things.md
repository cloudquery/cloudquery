# Table: aws_iot_things


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|principals|StringArray|
|arn (PK)|String|
|attributes|JSON|
|thing_name|String|
|thing_type_name|String|
|version|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|