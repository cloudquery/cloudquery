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
|trail|JSON|
|tags|JSON|