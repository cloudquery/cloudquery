# Table: aws_iot_things

https://docs.aws.amazon.com/iot/latest/apireference/API_ThingAttribute.html

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
|principals|StringArray|
|arn (PK)|String|
|attributes|JSON|
|thing_name|String|
|thing_type_name|String|
|version|Int|