# Table: aws_kafka_configurations

This table shows data for Kafka Configurations.

https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-configuration.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|kafka_versions|`list<item: utf8, nullable>`|
|latest_revision|`json`|
|name|`utf8`|
|state|`utf8`|