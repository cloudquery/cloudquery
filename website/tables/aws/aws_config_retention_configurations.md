# Table: aws_config_retention_configurations

This table shows data for Config Retention Configurations.

https://docs.aws.amazon.com/config/latest/APIReference/API_RetentionConfiguration.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|retention_period_in_days|`int64`|