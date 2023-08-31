# Table: aws_iam_group_policies

This table shows data for IAM Group Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html

The composite primary key for this table is (**account_id**, **group_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_groups](aws_iam_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|group_arn (PK)|`utf8`|
|policy_document|`json`|
|group_name|`utf8`|
|policy_name (PK)|`utf8`|
|result_metadata|`json`|

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


