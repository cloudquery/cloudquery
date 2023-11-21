# Table: aws_cloudtrail_trail_advanced_event_selectors

This table shows data for AWS CloudTrail Trail Advanced Event Selectors.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudtrail_trails](aws_cloudtrail_trails.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|trail_arn|`utf8`|
|field_selectors|`json`|
|name|`utf8`|