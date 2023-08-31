# Table: aws_mq_broker_users

This table shows data for Amazon MQ Broker Users.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id-users-username.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_mq_brokers](aws_mq_brokers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_arn|`utf8`|
|broker_id|`utf8`|
|console_access|`bool`|
|groups|`list<item: utf8, nullable>`|
|pending|`json`|
|replication_user|`bool`|
|username|`utf8`|
|result_metadata|`json`|