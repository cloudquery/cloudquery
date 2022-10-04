# Table: aws_cloudtrail_trail_event_selectors



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_cloudtrail_trails`](aws_cloudtrail_trails.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|trail_arn|String|
|data_resources|JSON|
|exclude_management_event_sources|StringArray|
|include_management_events|Bool|
|read_write_type|String|