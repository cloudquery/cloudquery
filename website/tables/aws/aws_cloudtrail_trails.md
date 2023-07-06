# Table: aws_cloudtrail_trails

This table shows data for AWS CloudTrail Trails.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Trail.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_cloudtrail_trails:
  - [aws_cloudtrail_trail_event_selectors](aws_cloudtrail_trail_event_selectors)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|cloudwatch_logs_log_group_name|`utf8`|
|arn (PK)|`utf8`|
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

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket

```sql
SELECT
  'Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket'
    AS title,
  t.account_id,
  t.arn AS resource_id,
  CASE
  WHEN b.logging_target_bucket IS NULL OR b.logging_target_prefix IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails AS t
  INNER JOIN aws_s3_buckets AS b ON t.s3_bucket_name = b.name;
```

### Ensure CloudTrail is enabled in all regions

```sql
SELECT
  'Ensure CloudTrail is enabled in all regions' AS title,
  aws_cloudtrail_trails.account_id,
  arn AS resource_id,
  CASE
  WHEN is_multi_region_trail = false
  OR (
      is_multi_region_trail = true
      AND (read_write_type != 'All' OR include_management_events = false)
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails
  INNER JOIN aws_cloudtrail_trail_event_selectors ON
      aws_cloudtrail_trails.arn = aws_cloudtrail_trail_event_selectors.trail_arn
      AND aws_cloudtrail_trails.region
        = aws_cloudtrail_trail_event_selectors.region
      AND aws_cloudtrail_trails.account_id
        = aws_cloudtrail_trail_event_selectors.account_id;
```

### CloudTrail trails should be integrated with CloudWatch Logs

```sql
SELECT
  'CloudTrail trails should be integrated with CloudWatch Logs' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN cloud_watch_logs_log_group_arn IS NULL
  OR (status->>'LatestCloudWatchLogsDeliveryTime')::TIMESTAMP
    < (now() - '1 days'::INTERVAL)
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails;
```

### Ensure CloudTrail log file validation is enabled

```sql
SELECT
  'Ensure CloudTrail log file validation is enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN log_file_validation_enabled = false THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails;
```

### CloudTrail should have encryption at rest enabled

```sql
SELECT
  'CloudTrail should have encryption at rest enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN kms_key_id IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_cloudtrail_trails;
```


