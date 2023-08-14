# Table: aws_iam_user_access_keys

This table shows data for IAM User Access Keys.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html

The composite primary key for this table is (**account_id**, **user_arn**, **access_key_id**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|access_key_id (PK)|`utf8`|
|user_id|`utf8`|
|last_used|`timestamp[us, tz=UTC]`|
|last_used_service_name|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|user_name|`utf8`|
|last_rotated|`timestamp[us, tz=UTC]`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### IAM users'' access keys should be rotated every 90 days or less

```sql
SELECT
  e'IAM users\' access keys should be rotated every 90 days or less' AS title,
  account_id,
  access_key_id AS resource_id,
  CASE
  WHEN date_part('day', now() - last_rotated) > 90 THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_user_access_keys;
```

### Unused IAM user credentials should be removed

```sql
SELECT
  'Unused IAM user credentials should be removed' AS title,
  account_id,
  access_key_id AS resource_id,
  CASE
  WHEN date_part('day', now() - last_used) > 90 THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_user_access_keys;
```

### Ensure access keys are rotated every 90 days or less

```sql
SELECT
  'Ensure access keys are rotated every 90 days or less' AS title,
  account_id,
  user_arn,
  CASE
  WHEN last_rotated < (now() - '90 days'::INTERVAL) THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_user_access_keys;
```

### Ensure no root account access key exists (Scored)

```sql
SELECT
  'Ensure no root account access key exists (Scored)' AS title,
  account_id,
  user_arn AS resource_id,
  CASE
  WHEN user_name IN ('<root>', '<root_account>') THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_user_access_keys;
```

### Ensure credentials unused for 90 days or greater are disabled (Scored)

```sql
SELECT
  'Ensure credentials unused for 90 days or greater are disabled (Scored)'
    AS title,
  split_part(r.arn, ':', 5) AS account_id,
  r.arn,
  CASE
  WHEN (
    r.password_status IN ('TRUE', 'true')
    AND r.password_last_used < (now() - '90 days'::INTERVAL)
    OR k.last_used < (now() - '90 days'::INTERVAL)
  )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_credential_reports AS r
  LEFT JOIN aws_iam_user_access_keys AS k ON k.user_arn = r.arn;
```

### Ensure credentials unused for 45 days or greater are disabled (Automated)

```sql
SELECT
  'Ensure credentials unused for 45 days or greater are disabled (Automated)'
    AS title,
  split_part(r.arn, ':', 5) AS account_id,
  r.arn,
  CASE
  WHEN (
    r.password_status IN ('TRUE', 'true')
    AND r.password_last_used < (now() - '45 days'::INTERVAL)
    OR (
        r.password_status IN ('TRUE', 'true')
        AND r.password_last_used IS NULL
        AND r.password_last_changed < (now() - '45 days'::INTERVAL)
      )
    OR k.last_used < (now() - '45 days'::INTERVAL)
  )
  OR (
      r.access_key1_active
      AND r.access_key_1_last_used_date < (now() - '45 days'::INTERVAL)
    )
  OR (
      r.access_key1_active
      AND r.access_key_1_last_used_date IS NULL
      AND access_key_1_last_rotated < (now() - '45 days'::INTERVAL)
    )
  OR (
      r.access_key2_active
      AND r.access_key_2_last_used_date < (now() - '45 days'::INTERVAL)
    )
  OR (
      r.access_key2_active
      AND r.access_key_2_last_used_date IS NULL
      AND access_key_2_last_rotated < (now() - '45 days'::INTERVAL)
    )
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_credential_reports AS r
  LEFT JOIN aws_iam_user_access_keys AS k ON k.user_arn = r.arn;
```

### Ensure there is only one active access key available for any single IAM user (Automated)

```sql
SELECT
  'Ensure there is only one active access key available for any single IAM user (Automated)'
    AS title,
  account_id,
  user_arn,
  CASE
  WHEN count(*) > 1 THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_user_access_keys
WHERE
  status = 'Active'
GROUP BY
  account_id, user_arn;
```


