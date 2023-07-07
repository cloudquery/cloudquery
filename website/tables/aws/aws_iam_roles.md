# Table: aws_iam_roles

This table shows data for IAM Roles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_attached_policies](aws_iam_role_attached_policies)
  - [aws_iam_role_last_accessed_details](aws_iam_role_last_accessed_details)
  - [aws_iam_role_policies](aws_iam_role_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|assume_role_policy_document|`json`|
|tags|`json`|
|arn (PK)|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|path|`utf8`|
|role_id|`utf8`|
|role_name|`utf8`|
|description|`utf8`|
|max_session_duration|`int64`|
|permissions_boundary|`json`|
|role_last_used|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


