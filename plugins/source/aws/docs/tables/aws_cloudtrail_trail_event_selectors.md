# Table: aws_cloudtrail_trail_event_selectors

This table shows data for AWS CloudTrail Trail Event Selectors.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetEventSelectors.html

The primary key for this table is **trail_arn**.

## Relations

This table depends on [aws_cloudtrail_trails](aws_cloudtrail_trails.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|advanced_event_selectors|`json`|
|event_selectors|`json`|
|trail_arn (PK)|`utf8`|