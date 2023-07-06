# Table: aws_dynamodb_table_continuous_backups

This table shows data for Amazon DynamoDB Table Continuous Backups.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_dynamodb_tables](aws_dynamodb_tables).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|table_arn|`utf8`|
|continuous_backups_status|`utf8`|
|point_in_time_recovery_description|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


