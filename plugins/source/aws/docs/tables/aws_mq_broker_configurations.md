# Table: aws_mq_broker_configurations

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_mq_brokers](aws_mq_brokers.md).

The following tables depend on aws_mq_broker_configurations:
  - [aws_mq_broker_configuration_revisions](aws_mq_broker_configuration_revisions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|broker_arn|String|
|arn|String|
|authentication_strategy|String|
|created|Timestamp|
|description|String|
|engine_type|String|
|engine_version|String|
|id|String|
|latest_revision|JSON|
|name|String|
|tags|JSON|