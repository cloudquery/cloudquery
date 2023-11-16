# Table: aws_quicksight_data_sources

This table shows data for QuickSight Data Sources.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSource.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|data_source_id|`utf8`|
|error_info|`json`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|secret_arn|`utf8`|
|ssl_properties|`json`|
|status|`utf8`|
|type|`utf8`|
|vpc_connection_properties|`json`|