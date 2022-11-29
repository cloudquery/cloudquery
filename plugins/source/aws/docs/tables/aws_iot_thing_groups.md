# Table: aws_iot_thing_groups



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
|things_in_group|StringArray|
|policies|StringArray|
|tags|JSON|
|arn (PK)|String|
|index_name|String|
|query_string|String|
|query_version|String|
|status|String|
|thing_group_id|String|
|thing_group_metadata|JSON|
|thing_group_name|String|
|thing_group_properties|JSON|
|version|Int|
|result_metadata|JSON|