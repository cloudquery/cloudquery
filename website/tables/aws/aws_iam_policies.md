# Table: aws_iam_policies

This table shows data for IAM Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_policies:
  - [aws_iam_policy_last_accessed_details](aws_iam_policy_last_accessed_details)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|tags|`json`|
|policy_version_list|`json`|
|arn|`utf8`|
|attachment_count|`int64`|
|create_date|`timestamp[us, tz=UTC]`|
|default_version_id|`utf8`|
|description|`utf8`|
|is_attachable|`bool`|
|path|`utf8`|
|permissions_boundary_usage_count|`int64`|
|policy_id|`utf8`|
|policy_name|`utf8`|
|update_date|`timestamp[us, tz=UTC]`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### IAM policies should not allow full ''*'' administrative privileges

```sql
WITH
  pvs
    AS (
      SELECT
        id, (v->>'Document')::JSONB AS document
      FROM
        aws_iam_policies,
        jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
    ),
  violations
    AS (
      SELECT
        id, count(*) AS violations
      FROM
        pvs,
        jsonb_array_elements(
          CASE jsonb_typeof(document->'Statement')
          WHEN 'string' THEN jsonb_build_array(document->>'Statement')
          WHEN 'array' THEN document->'Statement'
          END
        )
          AS statement,
        jsonb_array_elements_text(
          CASE jsonb_typeof(statement->'Resource')
          WHEN 'string' THEN jsonb_build_array(statement->>'Resource')
          WHEN 'array' THEN statement->'Resource'
          END
        )
          AS resource,
        jsonb_array_elements_text(
          CASE jsonb_typeof(statement->'Action')
          WHEN 'string' THEN jsonb_build_array(statement->>'Action')
          WHEN 'array' THEN statement->'Action'
          END
        )
          AS action
      WHERE
        statement->>'Effect' = 'Allow'
        AND resource = '*'
        AND (action = '*' OR action = '*:*')
      GROUP BY
        id
    )
SELECT
  DISTINCT
  e'IAM policies should not allow full \'*\' administrative privileges'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN violations.id IS NOT NULL AND violations.violations > 0 THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_iam_policies LEFT JOIN violations ON violations.id = aws_iam_policies.id;
```

### IAM policies should not allow full ''*'' administrative privileges

```sql
WITH
  iam_policies
    AS (
      SELECT
        id, (v->>'Document')::JSONB AS document
      FROM
        aws_iam_policies,
        jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
      WHERE
        aws_iam_policies.default_version_id = v->>'VersionId'
        AND arn NOT LIKE 'arn:aws:iam::aws:policy%'
    ),
  policy_statements
    AS (
      SELECT
        id,
        jsonb_array_elements(
          CASE jsonb_typeof(document->'Statement')
          WHEN 'string' THEN jsonb_build_array(document->>'Statement')
          WHEN 'array' THEN document->'Statement'
          END
        )
          AS statement
      FROM
        iam_policies
    ),
  allow_all_statements
    AS (
      SELECT
        id, count(statement) AS statements_count
      FROM
        policy_statements
      WHERE
        (statement->>'Action' = '*' OR statement->>'Action' LIKE '%"*"%')
        AND statement->>'Effect' = 'Allow'
        AND (
            statement->>'Resource' = '*'
            OR statement->>'Resource' LIKE '%"*"%'
          )
      GROUP BY
        id
    )
SELECT
  DISTINCT
  e'IAM policies should not allow full \'*\' administrative privileges'
    AS title,
  aws_iam_policies.account_id,
  aws_iam_policies.arn AS resource_id,
  CASE WHEN statements_count > 0 THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_iam_policies
  LEFT JOIN allow_all_statements ON
      aws_iam_policies.id = allow_all_statements.id;
```

### IAM customer managed policies that you create should not allow wildcard actions for services

```sql
WITH
  policy_statements
    AS (
      SELECT
        aws_iam_policies.id,
        jsonb_array_elements(
          CASE jsonb_typeof((v->>'Document')::JSONB->'Statement')
          WHEN 'string'
          THEN jsonb_build_array((v->>'Document')::JSONB->>'Statement')
          WHEN 'array' THEN (v->>'Document')::JSONB->'Statement'
          END
        )
          AS statement
      FROM
        aws_iam_policies,
        jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
      WHERE
        aws_iam_policies.arn NOT LIKE 'arn:aws:iam::aws:policy%'
    ),
  allow_all_statements
    AS (
      SELECT
        id, count(statement) AS statements_count
      FROM
        policy_statements
      WHERE
        statement->>'Effect' = 'Allow'
        AND (
            statement->>'Action' LIKE '%*%'
            OR statement->>'NotAction' LIKE '%*%'
          )
      GROUP BY
        id
    )
SELECT
  DISTINCT
  'IAM customer managed policies that you create should not allow wildcard actions for services'
    AS title,
  aws_iam_policies.account_id,
  aws_iam_policies.arn AS resource_id,
  CASE WHEN statements_count > 0 THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_iam_policies
  LEFT JOIN allow_all_statements ON
      aws_iam_policies.id = allow_all_statements.id;
```

### IAM customer managed policies should not allow decryption and re-encryption actions on all KMS keys

```sql
WITH
  iam_policies
    AS (
      SELECT
        (v->>'Document')::JSONB AS document, account_id, arn, id
      FROM
        aws_iam_policies,
        jsonb_array_elements(aws_iam_policies.policy_version_list) AS v
    ),
  violations
    AS (
      SELECT
        DISTINCT id
      FROM
        iam_policies,
        jsonb_array_elements(
          CASE jsonb_typeof(document->'Statement')
          WHEN 'string' THEN jsonb_build_array(document->>'Statement')
          WHEN 'array' THEN document->'Statement'
          END
        )
          AS statement
      WHERE
        NOT
          (
            arn LIKE 'arn:aws:iam::aws:policy%'
            OR arn LIKE 'arn:aws-us-gov:iam::aws:policy%'
          )
        AND statement->>'Effect' = 'Allow'
        AND statement->'Resource'
          ?| ARRAY[
              '*',
              'arn:aws:kms:*:' || account_id || ':key/*',
              'arn:aws:kms:*:' || account_id || ':alias/*'
            ]
        AND statement->'Action'
          ?| ARRAY[
              '*',
              'kms:*',
              'kms:decrypt',
              'kms:reencryptfrom',
              'kms:reencrypt*'
            ]
    )
SELECT
  'IAM customer managed policies should not allow decryption and re-encryption actions on all KMS keys'
    AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN violations.id IS NOT NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_iam_policies LEFT JOIN violations ON violations.id = aws_iam_policies.id;
```


