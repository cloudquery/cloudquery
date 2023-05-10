# Table: aws_config_configuration_aggregators

This table shows data for Config Configuration Aggregators.

https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationAggregator.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|account_aggregation_sources|JSON|
|configuration_aggregator_arn|String|
|configuration_aggregator_name|String|
|created_by|String|
|creation_time|Timestamp|
|last_updated_time|Timestamp|
|organization_aggregation_source|JSON|