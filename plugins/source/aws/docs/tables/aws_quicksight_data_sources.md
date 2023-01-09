# Table: aws_quicksight_data_sources

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSource.html

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
|tags|JSON|
|arn (PK)|String|
|created_time|Timestamp|
|data_source_id|String|
|error_info|JSON|
|last_updated_time|Timestamp|
|name|String|
|secret_arn|String|
|ssl_properties|JSON|
|status|String|
|type|String|
|vpc_connection_properties|JSON|