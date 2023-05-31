# Table: aws_cloudtrail_trail_event_selectors

This table shows data for AWS CloudTrail Trail Event Selectors.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_EventSelector.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudtrail_trails](aws_cloudtrail_trails).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|trail_arn|`utf8`|
|data_resources|`json`|
|exclude_management_event_sources|`list<item: utf8, nullable>`|
|include_management_events|`bool`|
|read_write_type|`utf8`|