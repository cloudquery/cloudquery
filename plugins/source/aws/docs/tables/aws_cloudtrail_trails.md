# Table: aws_cloudtrail_trails

This table shows data for AWS CloudTrail Trails.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Trail.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_cloudtrail_trails:
  - [aws_cloudtrail_trail_event_selectors](aws_cloudtrail_trail_event_selectors.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cloudwatch_logs_log_group_name|`utf8`|
|arn|`utf8`|
|status|`json`|
|cloud_watch_logs_log_group_arn|`utf8`|
|cloud_watch_logs_role_arn|`utf8`|
|has_custom_event_selectors|`bool`|
|has_insight_selectors|`bool`|
|home_region|`utf8`|
|include_global_service_events|`bool`|
|is_multi_region_trail|`bool`|
|is_organization_trail|`bool`|
|kms_key_id|`utf8`|
|log_file_validation_enabled|`bool`|
|name|`utf8`|
|s3_bucket_name|`utf8`|
|s3_key_prefix|`utf8`|
|sns_topic_arn|`utf8`|
|sns_topic_name|`utf8`|
|trail_arn|`utf8`|
|tags|`json`|