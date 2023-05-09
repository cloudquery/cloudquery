# Table: aws_config_retention_configurations

This table shows data for Config Retention Configurations.

https://docs.aws.amazon.com/config/latest/APIReference/API_RetentionConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|retention_period_in_days|Int|