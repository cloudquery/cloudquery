# Table: aws_dynamodb_tables

This table shows data for Amazon DynamoDB Tables.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_dynamodb_tables:
  - [aws_dynamodb_table_continuous_backups](aws_dynamodb_table_continuous_backups)
  - [aws_dynamodb_table_replica_auto_scalings](aws_dynamodb_table_replica_auto_scalings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|archival_summary|`json`|
|attribute_definitions|`json`|
|billing_mode_summary|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|deletion_protection_enabled|`bool`|
|global_secondary_indexes|`json`|
|global_table_version|`utf8`|
|item_count|`int64`|
|key_schema|`json`|
|latest_stream_arn|`utf8`|
|latest_stream_label|`utf8`|
|local_secondary_indexes|`json`|
|provisioned_throughput|`json`|
|replicas|`json`|
|restore_summary|`json`|
|sse_description|`json`|
|stream_specification|`json`|
|table_arn|`utf8`|
|table_class_summary|`json`|
|table_id|`utf8`|
|table_name|`utf8`|
|table_size_bytes|`int64`|
|table_status|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### DynamoDB tables should automatically scale capacity with demand

```sql
SELECT
  'DynamoDB tables should automatically scale capacity with demand' AS title,
  t.account_id,
  pr.arn AS resource_id,
  CASE
  WHEN t.billing_mode_summary->>'BillingMode' IS DISTINCT FROM 'PAY_PER_REQUEST'
  AND (
      (
        s.replica_provisioned_read_capacity_auto_scaling_settings->>'AutoScalingDisabled'
      )::BOOL
      IS NOT false
      OR (
          s.replica_provisioned_write_capacity_auto_scaling_settings->>'AutoScalingDisabled'
        )::BOOL
        IS NOT false
    )
  AND (pr._cq_id IS NULL OR pw._cq_id IS NULL)
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_dynamodb_tables AS t
  LEFT JOIN aws_dynamodb_table_replica_auto_scalings AS s ON s.table_arn = t.arn
  LEFT JOIN aws_applicationautoscaling_policies AS pr ON
      pr.service_namespace = 'dynamodb'
      AND pr.resource_id = concat('table/', t.table_name)
      AND pr.policy_type = 'TargetTrackingScaling'
      AND pr.scalable_dimension = 'dynamodb:table:ReadCapacityUnits'
  LEFT JOIN aws_applicationautoscaling_policies AS pw ON
      pw.service_namespace = 'dynamodb'
      AND pw.resource_id = concat('table/', t.table_name)
      AND pw.policy_type = 'TargetTrackingScaling'
      AND pw.scalable_dimension = 'dynamodb:table:WriteCapacityUnits';
```

### DynamoDB tables should have point-in-time recovery enabled

```sql
SELECT
  'DynamoDB tables should have point-in-time recovery enabled' AS title,
  t.account_id,
  t.arn AS resource_id,
  CASE
  WHEN b.point_in_time_recovery_description->>'PointInTimeRecoveryStatus'
  IS DISTINCT FROM 'ENABLED'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_dynamodb_tables AS t
  LEFT JOIN aws_dynamodb_table_continuous_backups AS b ON b.table_arn = t.arn;
```

### DynamoDB table with no items

```sql
SELECT
  'DynamoDB table with no items' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_dynamodb_tables
WHERE
  item_count = 0;
```


