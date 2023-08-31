# Table: aws_mq_broker_configuration_revisions

This table shows data for Amazon MQ Broker Configuration Revisions.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_mq_broker_configurations](aws_mq_broker_configurations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_configuration_arn|`utf8`|
|data|`json`|
|configuration_id|`utf8`|
|created|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|result_metadata|`json`|