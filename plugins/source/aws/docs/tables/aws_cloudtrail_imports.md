# Table: aws_cloudtrail_imports

This table shows data for AWS CloudTrail Imports.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetImport.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|destinations|`list<item: utf8, nullable>`|
|end_event_time|`timestamp[us, tz=UTC]`|
|import_id|`utf8`|
|import_source|`json`|
|import_statistics|`json`|
|import_status|`utf8`|
|start_event_time|`timestamp[us, tz=UTC]`|
|updated_timestamp|`timestamp[us, tz=UTC]`|