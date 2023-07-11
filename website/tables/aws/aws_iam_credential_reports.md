# Table: aws_iam_credential_reports

This table shows data for IAM Credential Reports.

https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html#id_credentials_understanding_the_report_format

The composite primary key for this table is (**arn**, **user_creation_time**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|arn (PK)|`utf8`|
|user_creation_time (PK)|`timestamp[us, tz=UTC]`|
|password_last_changed|`timestamp[us, tz=UTC]`|
|password_next_rotation|`timestamp[us, tz=UTC]`|
|access_key_1_last_rotated|`timestamp[us, tz=UTC]`|
|access_key_2_last_rotated|`timestamp[us, tz=UTC]`|
|cert_1_last_rotated|`timestamp[us, tz=UTC]`|
|cert_2_last_rotated|`timestamp[us, tz=UTC]`|
|access_key_1_last_used_date|`timestamp[us, tz=UTC]`|
|access_key_2_last_used_date|`timestamp[us, tz=UTC]`|
|password_last_used|`timestamp[us, tz=UTC]`|
|password_enabled|`utf8`|
|user|`utf8`|
|password_status|`utf8`|
|mfa_active|`bool`|
|access_key1_active|`bool`|
|access_key2_active|`bool`|
|cert1_active|`bool`|
|cert2_active|`bool`|
|access_key1_last_used_region|`utf8`|
|access_key1_last_used_service|`utf8`|
|access_key2_last_used_region|`utf8`|
|access_key2_last_used_service|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure hardware MFA is enabled for the "root" account (Scored)

```sql
SELECT
  'Ensure hardware MFA is enabled for the "root" account (Scored)' AS title,
  split_part(cr.arn, ':', 5) AS account_id,
  cr.arn AS resource_id,
  CASE
  WHEN mfa.serial_number IS NULL OR cr.mfa_active = false THEN 'fail'
  WHEN mfa.serial_number IS NOT NULL AND cr.mfa_active = true THEN 'pass'
  END
    AS status
FROM
  aws_iam_credential_reports AS cr
  LEFT JOIN aws_iam_virtual_mfa_devices AS mfa ON mfa.user->>'Arn' = cr.arn
WHERE
  cr.user = '<root_account>'
GROUP BY
  mfa.serial_number, cr.mfa_active, cr.arn;
```

### Ensure MFA is enabled for all IAM users that have a console password (Scored)

```sql
SELECT
  'Ensure MFA is enabled for all IAM users that have a console password (Scored)'
    AS title,
  split_part(arn, ':', 5) AS account_id,
  arn AS resource_id,
  CASE
  WHEN password_status IN ('TRUE', 'true') AND NOT mfa_active THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_credential_reports;
```

### Ensure MFA is enabled for the "root" account

```sql
SELECT
  'Ensure MFA is enabled for the "root" account' AS title,
  split_part(arn, ':', 5) AS account_id,
  arn AS resource_id,
  CASE
  WHEN current_user() = '<root_account>' AND NOT mfa_active THEN 'fail'
  WHEN current_user() = '<root_account>' AND mfa_active THEN 'pass'
  END
    AS status
FROM
  aws_iam_credential_reports;
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


