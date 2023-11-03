# Table: aws_applicationautoscaling_policies

This table shows data for Application Auto Scaling Policies.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|policy_arn|`utf8`|
|policy_name|`utf8`|
|policy_type|`utf8`|
|resource_id|`utf8`|
|scalable_dimension|`utf8`|
|service_namespace|`utf8`|
|alarms|`json`|
|step_scaling_policy_configuration|`json`|
|target_tracking_scaling_policy_configuration|`json`|

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


