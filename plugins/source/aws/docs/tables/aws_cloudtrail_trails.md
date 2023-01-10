# Table: aws_cloudtrail_trails

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Trail.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_cloudtrail_trails:
  - [aws_cloudtrail_trail_event_selectors](aws_cloudtrail_trail_event_selectors.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cloudwatch_logs_log_group_name|String|
|arn (PK)|String|
|status|JSON|
|cloud_watch_logs_log_group_arn|String|
|cloud_watch_logs_role_arn|String|
|has_custom_event_selectors|Bool|
|has_insight_selectors|Bool|
|home_region|String|
|include_global_service_events|Bool|
|is_multi_region_trail|Bool|
|is_organization_trail|Bool|
|kms_key_id|String|
|log_file_validation_enabled|Bool|
|name|String|
|s3_bucket_name|String|
|s3_key_prefix|String|
|sns_topic_arn|String|
|sns_topic_name|String|
|trail_arn|String|
|tags|JSON|