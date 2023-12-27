# Table: aws_mq_broker_configuration_revisions

This table shows data for Amazon MQ Broker Configuration Revisions.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html

The composite primary key for this table is (**broker_configuration_arn**, **revision**, **configuration_id**).

## Relations

This table depends on [aws_mq_broker_configurations](aws_mq_broker_configurations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_configuration_arn (PK)|`utf8`|
|revision (PK)|`int32`|
|data|`json`|
|configuration_id (PK)|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|description|`utf8`|