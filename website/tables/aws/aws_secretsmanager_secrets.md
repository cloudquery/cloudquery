# Table: aws_secretsmanager_secrets

This table shows data for AWS Secrets Manager Secrets.

https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_secretsmanager_secrets:
  - [aws_secretsmanager_secret_versions](aws_secretsmanager_secret_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|policy|`json`|
|tags|`json`|
|created_date|`timestamp[us, tz=UTC]`|
|deleted_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|kms_key_id|`utf8`|
|last_accessed_date|`timestamp[us, tz=UTC]`|
|last_changed_date|`timestamp[us, tz=UTC]`|
|last_rotated_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|next_rotation_date|`timestamp[us, tz=UTC]`|
|owning_service|`utf8`|
|primary_region|`utf8`|
|replication_status|`json`|
|rotation_enabled|`bool`|
|rotation_lambda_arn|`utf8`|
|rotation_rules|`json`|
|version_ids_to_stages|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Remove unused Secrets Manager secrets

```sql
SELECT
  'Remove unused Secrets Manager secrets' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (
    last_accessed_date IS NULL
    AND created_date > now() - '90 days'::INTERVAL
  )
  OR (
      last_accessed_date IS NOT NULL
      AND last_accessed_date > now() - '90 days'::INTERVAL
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_secretsmanager_secrets;
```

### Secrets Manager secrets configured with automatic rotation should rotate successfully

```sql
SELECT
  'Secrets Manager secrets configured with automatic rotation should rotate successfully'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (
    last_rotated_date IS NULL
    AND created_date
      > now()
        - '1 day'::INTERVAL * (rotation_rules->>'AutomaticallyAfterDays')::INT8
  )
  OR (
      last_rotated_date IS NOT NULL
      AND last_rotated_date
        > now()
          - '1 day'::INTERVAL
            * (rotation_rules->>'AutomaticallyAfterDays')::INT8
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_secretsmanager_secrets;
```

### Secrets Manager secrets should be rotated within a specified number of days

```sql
SELECT
  'Secrets Manager secrets should be rotated within a specified number of days'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (
    last_rotated_date IS NULL
    AND created_date > now() - '90 days'::INTERVAL
  )
  OR (
      last_rotated_date IS NOT NULL
      AND last_rotated_date > now() - '90 days'::INTERVAL
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_secretsmanager_secrets;
```

### Secrets Manager secrets should have automatic rotation enabled

```sql
SELECT
  'Secrets Manager secrets should have automatic rotation enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN rotation_enabled IS NOT true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_secretsmanager_secrets;
```


