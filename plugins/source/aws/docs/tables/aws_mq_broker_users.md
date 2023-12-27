# Table: aws_mq_broker_users

This table shows data for Amazon MQ Broker Users.

https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id-users-username.html

The composite primary key for this table is (**broker_arn**, **username**).

## Relations

This table depends on [aws_mq_brokers](aws_mq_brokers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|broker_arn (PK)|`utf8`|
|broker_id|`utf8`|
|console_access|`bool`|
|groups|`list<item: utf8, nullable>`|
|pending|`json`|
|replication_user|`bool`|
|username (PK)|`utf8`|