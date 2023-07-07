# Table: aws_iam_password_policies

This table shows data for IAM Password Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_PasswordPolicy.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|allow_users_to_change_password|`bool`|
|expire_passwords|`bool`|
|hard_expiry|`bool`|
|max_password_age|`int64`|
|minimum_password_length|`int64`|
|password_reuse_prevention|`int64`|
|require_lowercase_characters|`bool`|
|require_numbers|`bool`|
|require_symbols|`bool`|
|require_uppercase_characters|`bool`|
|policy_exists|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure IAM password policy expires passwords within 90 days or less

```sql
SELECT
  'Ensure IAM password policy expires passwords within 90 days or less'
    AS title,
  account_id,
  account_id,
  CASE
  WHEN (max_password_age IS NULL OR max_password_age > 90)
  OR policy_exists = false
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy requires minimum length of 14 or greater

```sql
SELECT
  'Ensure IAM password policy requires minimum length of 14 or greater'
    AS title,
  account_id,
  account_id,
  CASE
  WHEN minimum_password_length < 14 OR policy_exists = false THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy requires at least one lowercase letter

```sql
SELECT
  'Ensure IAM password policy requires at least one lowercase letter' AS title,
  account_id,
  account_id,
  CASE
  WHEN require_lowercase_characters = false OR policy_exists = false THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy requires at least one number

```sql
SELECT
  'Ensure IAM password policy requires at least one number' AS title,
  account_id,
  account_id,
  CASE
  WHEN require_numbers = false OR policy_exists = false THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy requires at least one symbol

```sql
SELECT
  'Ensure IAM password policy requires at least one symbol' AS title,
  account_id,
  account_id,
  CASE
  WHEN require_symbols = false OR policy_exists = false THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy requires at least one uppercase letter

```sql
SELECT
  'Ensure IAM password policy requires at least one uppercase letter' AS title,
  account_id,
  account_id,
  CASE
  WHEN require_uppercase_characters IS NOT true OR policy_exists IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_password_policies;
```

### Ensure IAM password policy prevents password reuse

```sql
SELECT
  'Ensure IAM password policy prevents password reuse' AS title,
  account_id,
  account_id,
  CASE
  WHEN (password_reuse_prevention IS NULL OR password_reuse_prevention > 24)
  OR policy_exists = false
  THEN 'fail'
  ELSE 'pass'
  END
FROM
  aws_iam_password_policies;
```

### Password policies for IAM users should have strong configurations

```sql
SELECT
  'Password policies for IAM users should have strong configurations' AS title,
  account_id,
  account_id AS resource_id,
  CASE
  WHEN (
    require_uppercase_characters IS NOT true
    OR require_lowercase_characters IS NOT true
    OR require_numbers IS NOT true
    OR minimum_password_length < 14
    OR password_reuse_prevention IS NULL
    OR max_password_age IS NULL
    OR policy_exists IS NOT true
  )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_password_policies;
```


