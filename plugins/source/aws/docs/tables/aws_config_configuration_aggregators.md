# Table: aws_config_configuration_aggregators

This table shows data for Config Configuration Aggregators.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationAggregator.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|account_aggregation_sources|`json`|
|configuration_aggregator_arn|`utf8`|
|configuration_aggregator_name|`utf8`|
|created_by|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|organization_aggregation_source|`json`|