# Table: aws_mq_broker_configurations

This table shows data for Amazon MQ Broker Configurations.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_mq_brokers](aws_mq_brokers).

The following tables depend on aws_mq_broker_configurations:
  - [aws_mq_broker_configuration_revisions](aws_mq_broker_configuration_revisions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|broker_arn|String|
|arn (PK)|String|
|authentication_strategy|String|
|created|Timestamp|
|description|String|
|engine_type|String|
|engine_version|String|
|id|String|
|latest_revision|JSON|
|name|String|
|tags|JSON|