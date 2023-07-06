# Table: aws_iam_users

This table shows data for IAM Users.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_iam_users:
  - [aws_iam_signing_certificates](aws_iam_signing_certificates)
  - [aws_iam_ssh_public_keys](aws_iam_ssh_public_keys)
  - [aws_iam_user_access_keys](aws_iam_user_access_keys)
  - [aws_iam_user_attached_policies](aws_iam_user_attached_policies)
  - [aws_iam_user_groups](aws_iam_user_groups)
  - [aws_iam_user_last_accessed_details](aws_iam_user_last_accessed_details)
  - [aws_iam_user_policies](aws_iam_user_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|create_date|`timestamp[us, tz=UTC]`|
|path|`utf8`|
|user_id|`utf8`|
|user_name|`utf8`|
|password_last_used|`timestamp[us, tz=UTC]`|
|permissions_boundary|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Avoid the use of "root" account. Show used in last 30 days (Scored)

```sql
SELECT
  'Avoid the use of "root" account. Show used in last 30 days (Scored)'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN user_name = '<root_account>'
  AND password_last_used > (now() - '30 days'::INTERVAL)
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_users;
```

### IAM users should not have IAM policies attached

```sql
SELECT
  DISTINCT
  'IAM users should not have IAM policies attached' AS title,
  aws_iam_users.account_id,
  arn AS resource_id,
  CASE
  WHEN aws_iam_user_attached_policies.user_arn IS NOT NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_users
  LEFT JOIN aws_iam_user_attached_policies ON
      aws_iam_users.arn = aws_iam_user_attached_policies.user_arn;
```

### IAM principals should not have IAM inline policies that allow decryption and re-encryption actions on all KMS keys

```sql
SELECT
  'IAM principals should not have IAM inline policies that allow decryption and re-encryption actions on all KMS keys'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  (
    SELECT
      statement, aws_iam_users.account_id, arn, policy_name
    FROM
      aws_iam_user_policies
      CROSS JOIN LATERAL jsonb_array_elements(
          CASE jsonb_typeof(policy_document->'Statement')
          WHEN 'string' THEN jsonb_build_array(policy_document->>'Statement')
          WHEN 'array' THEN policy_document->'Statement'
          END
        )
          AS statement
      INNER JOIN aws_iam_users ON
          aws_iam_users.account_id = aws_iam_user_policies.account_id
          AND aws_iam_users.arn = aws_iam_user_policies.user_arn
    UNION
      SELECT
        statement, aws_iam_roles.account_id, arn, policy_name
      FROM
        aws_iam_role_policies
        CROSS JOIN LATERAL jsonb_array_elements(
            CASE jsonb_typeof(policy_document->'Statement')
            WHEN 'string' THEN jsonb_build_array(policy_document->>'Statement')
            WHEN 'array' THEN policy_document->'Statement'
            END
          )
            AS statement
        INNER JOIN aws_iam_roles ON
            aws_iam_roles.account_id = aws_iam_role_policies.account_id
            AND aws_iam_roles.arn = aws_iam_role_policies.role_arn
      WHERE
        lower(arn) NOT LIKE 'arn:aws:iam::%:role/aws-service-role/%'
    UNION
      SELECT
        statement, aws_iam_groups.account_id, arn, policy_name
      FROM
        aws_iam_group_policies
        CROSS JOIN LATERAL jsonb_array_elements(
            CASE jsonb_typeof(policy_document->'Statement')
            WHEN 'string' THEN jsonb_build_array(policy_document->>'Statement')
            WHEN 'array' THEN policy_document->'Statement'
            END
          )
            AS statement
        INNER JOIN aws_iam_groups ON
            aws_iam_groups.account_id = aws_iam_group_policies.account_id
            AND aws_iam_groups.arn = aws_iam_group_policies.group_arn
  )
    AS t
WHERE
  statement->>'Effect' = 'Allow'
  AND lower(statement::STRING)::JSONB->'resource'
    ?| ARRAY[
        '*',
        'arn:aws:kms:*:*:key/*',
        'arn:aws:kms:*:' || account_id || ':key/*arn:aws:kms:*:*:alias/*',
        'arn:aws:kms:*:' || account_id || ':alias/*'
      ]
  AND lower(statement::STRING)::JSONB->'action'
    ?| ARRAY['*', 'kms:*', 'kms:decrypt', 'kms:encrypt*', 'kms:reencryptfrom'];
```


