# Table: aws_quicksight_data_sources

This table shows data for QuickSight Data Sources.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSource.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
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