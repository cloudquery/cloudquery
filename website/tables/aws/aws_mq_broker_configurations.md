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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_arn|`utf8`|
|arn (PK)|`utf8`|
|authentication_strategy|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|engine_type|`utf8`|
|engine_version|`utf8`|
|id|`utf8`|
|latest_revision|`json`|
|name|`utf8`|
|tags|`json`|