# Table: aws_mq_broker_configuration_revisions



The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|broker_configuration_arn|String|
|data|JSON|
|configuration_id|String|
|created|Timestamp|
|description|String|
|result_metadata|JSON|