# Table: aws_appflow_flows

This table shows data for Amazon AppFlow Flows.

https://docs.aws.amazon.com/appflow/1.0/APIReference/API_DescribeFlow.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|description|`utf8`|
|destination_flow_config_list|`json`|
|flow_arn|`utf8`|
|flow_name|`utf8`|
|flow_status|`utf8`|
|flow_status_message|`utf8`|
|kms_arn|`utf8`|
|last_run_execution_details|`json`|
|last_run_metadata_catalog_details|`json`|
|last_updated_at|`timestamp[us, tz=UTC]`|
|last_updated_by|`utf8`|
|metadata_catalog_config|`json`|
|schema_version|`int64`|
|source_flow_config|`json`|
|tags|`json`|
|tasks|`json`|
|trigger_config|`json`|