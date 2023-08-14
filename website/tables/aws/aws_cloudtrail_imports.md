# Table: aws_cloudtrail_imports

This table shows data for AWS CloudTrail Imports.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetImport.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|created_timestamp|`timestamp[us, tz=UTC]`|
|destinations|`list<item: utf8, nullable>`|
|end_event_time|`timestamp[us, tz=UTC]`|
|import_id|`utf8`|
|import_source|`json`|
|import_statistics|`json`|
|import_status|`utf8`|
|start_event_time|`timestamp[us, tz=UTC]`|
|updated_timestamp|`timestamp[us, tz=UTC]`|