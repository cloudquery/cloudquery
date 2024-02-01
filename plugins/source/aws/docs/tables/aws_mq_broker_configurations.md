# Table: aws_mq_broker_configurations

This table shows data for Amazon MQ Broker Configurations.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_mq_brokers](aws_mq_brokers.md).

The following tables depend on aws_mq_broker_configurations:
  - [aws_mq_broker_configuration_revisions](aws_mq_broker_configuration_revisions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_arn|`utf8`|
|arn|`utf8`|
|authentication_strategy|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|engine_type|`utf8`|
|engine_version|`utf8`|
|id|`utf8`|
|latest_revision|`json`|
|name|`utf8`|
|tags|`json`|