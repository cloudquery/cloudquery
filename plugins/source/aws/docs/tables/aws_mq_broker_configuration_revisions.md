# Table: aws_mq_broker_configuration_revisions

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_mq_broker_configurations](aws_mq_broker_configurations.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|broker_configuration_arn|String|
|data|JSON|
|configuration_id|String|
|created|Timestamp|
|description|String|
|result_metadata|JSON|