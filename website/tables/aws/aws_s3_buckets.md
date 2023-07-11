# Table: aws_s3_buckets

This table shows data for S3 Buckets.

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_s3_buckets:
  - [aws_s3_bucket_cors_rules](aws_s3_bucket_cors_rules)
  - [aws_s3_bucket_encryption_rules](aws_s3_bucket_encryption_rules)
  - [aws_s3_bucket_grants](aws_s3_bucket_grants)
  - [aws_s3_bucket_lifecycles](aws_s3_bucket_lifecycles)
  - [aws_s3_bucket_websites](aws_s3_bucket_websites)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|replication_role|`utf8`|
|replication_rules|`json`|
|region|`utf8`|
|logging_target_bucket|`utf8`|
|logging_target_prefix|`utf8`|
|policy|`json`|
|policy_status|`json`|
|versioning_status|`utf8`|
|versioning_mfa_delete|`utf8`|
|block_public_acls|`bool`|
|block_public_policy|`bool`|
|ignore_public_acls|`bool`|
|restrict_public_buckets|`bool`|
|tags|`json`|
|ownership_controls|`list<item: utf8, nullable>`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket

```sql
SELECT
  'Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket'
    AS title,
  t.account_id,
  t.arn AS resource_id,
  CASE
  WHEN b.logging_target_bucket IS NULL OR b.logging_target_prefix IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails AS t
  INNER JOIN aws_s3_buckets AS b ON t.s3_bucket_name = b.name;
```

### S3 Block Public Access (bucket) setting should be enabled

```sql
SELECT
  'S3 Block Public Access (bucket) setting should be enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN block_public_acls IS NOT true
  OR block_public_policy IS NOT true
  OR ignore_public_acls IS NOT true
  OR restrict_public_buckets IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_s3_buckets;
```

### S3 buckets should deny non-HTTPS requests

```sql
SELECT
  'S3 buckets should deny non-HTTPS requests' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_s3_buckets
WHERE
  arn
  NOT IN (
      SELECT
        arn
      FROM
        (
          SELECT
            aws_s3_buckets.arn,
            statements,
            statements->'Principal' AS principals
          FROM
            aws_s3_buckets,
            jsonb_array_elements(
              CASE jsonb_typeof(policy->'Statement')
              WHEN 'string' THEN jsonb_build_array(policy->>'Statement')
              WHEN 'array' THEN policy->'Statement'
              END
            )
              AS statements
          WHERE
            statements->'Effect' = '"Deny"'
        )
          AS foo,
        jsonb_array_elements_text(
          statements->'Condition'->'Bool'->'aws:securetransport'
        )
          AS ssl
      WHERE
        principals = '"*"'
        OR (
            principals::JSONB ? 'AWS'
            AND (principals->'AWS' = '"*"' OR principals->'AWS' @> '"*"')
          )
          AND ssl::BOOL = false
    );
```

### Ensure MFA Delete is enabled on S3 buckets (Automated)

```sql
SELECT
  'Ensure MFA Delete is enabled on S3 buckets (Automated)' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN versioning_status IS DISTINCT FROM 'Enabled'
  OR versioning_mfa_delete IS DISTINCT FROM 'Enabled'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_s3_buckets;
```

### S3 buckets should prohibit public read access

```sql
WITH
  policy_allow_public
    AS (
      SELECT
        arn, count(*) AS statement_count
      FROM
        (
          SELECT
            aws_s3_buckets.arn, statements->'Principal' AS principals
          FROM
            aws_s3_buckets,
            jsonb_array_elements(
              CASE jsonb_typeof(policy::JSONB->'Statement')
              WHEN 'string' THEN jsonb_build_array(policy::JSONB->>'Statement')
              WHEN 'array' THEN policy::JSONB->'Statement'
              END
            )
              AS statements
          WHERE
            statements->'Effect' = '"Allow"'
        )
          AS foo
      WHERE
        principals = '"*"'
        OR (
            principals::JSONB ? 'AWS'
            AND (principals->'AWS' = '"*"' OR principals->'AWS' @> '"*"')
          )
      GROUP BY
        arn
    )
SELECT
  'S3 buckets should prohibit public read access' AS title,
  aws_s3_buckets.account_id,
  aws_s3_buckets.arn AS resource_id,
  'fail' AS status
FROM
  aws_s3_buckets
  LEFT JOIN aws_s3_bucket_grants ON
      aws_s3_buckets.arn = aws_s3_bucket_grants.bucket_arn
  LEFT JOIN policy_allow_public ON aws_s3_buckets.arn = policy_allow_public.arn
WHERE
  (
    aws_s3_buckets.block_public_acls != true
    AND (
        grantee->>'URI' = 'http://acs.amazonaws.com/groups/global/AllUsers'
        AND permission IN ('READ_ACP', 'FULL_CONTROL')
      )
  )
  OR (
      aws_s3_buckets.block_public_policy != true
      AND policy_allow_public.statement_count > 0
    );
```

### S3 buckets should prohibit public write access

```sql
WITH
  policy_allow_public
    AS (
      SELECT
        arn, count(*) AS statement_count
      FROM
        (
          SELECT
            aws_s3_buckets.arn, statements->'Principal' AS principals
          FROM
            aws_s3_buckets,
            jsonb_array_elements(
              CASE jsonb_typeof(policy::JSONB->'Statement')
              WHEN 'string' THEN jsonb_build_array(policy::JSONB->>'Statement')
              WHEN 'array' THEN policy::JSONB->'Statement'
              END
            )
              AS statements
          WHERE
            statements->'Effect' = '"Allow"'
        )
          AS foo
      WHERE
        principals = '"*"'
        OR (
            principals::JSONB ? 'AWS'
            AND (principals->'AWS' = '"*"' OR principals->'AWS' @> '"*"')
          )
      GROUP BY
        arn
    )
SELECT
  'S3 buckets should prohibit public write access' AS title,
  aws_s3_buckets.account_id,
  aws_s3_buckets.arn AS resource_id,
  'fail' AS status
FROM
  aws_s3_buckets
  LEFT JOIN aws_s3_bucket_grants ON
      aws_s3_buckets.arn = aws_s3_bucket_grants.bucket_arn
  LEFT JOIN policy_allow_public ON aws_s3_buckets.arn = policy_allow_public.arn
WHERE
  (
    aws_s3_buckets.block_public_acls != true
    AND (
        grantee->>'URI' = 'http://acs.amazonaws.com/groups/global/AllUsers'
        AND permission IN ('WRITE_ACP', 'FULL_CONTROL')
      )
  )
  OR (
      aws_s3_buckets.block_public_policy != true
      AND policy_allow_public.statement_count > 0
    );
```

### Amazon S3 permissions granted to other AWS accounts in bucket policies should be restricted

```sql
SELECT
  'Amazon S3 permissions granted to other AWS accounts in bucket policies should be restricted'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  (
    SELECT
      aws_s3_buckets.arn,
      account_id,
      name,
      region,
      CASE
      WHEN jsonb_typeof(statements->'Principal') = 'string'
      THEN jsonb_build_array(statements->'Principal')
      WHEN jsonb_typeof(statements->'Principal'->'AWS') = 'string'
      THEN jsonb_build_array(statements->'Principal'->'AWS')
      WHEN jsonb_typeof(statements->'Principal'->'AWS') = 'array'
      THEN statements->'Principal'->'AWS'
      END
        AS principals,
      CASE
      WHEN jsonb_typeof(statements->'Action') = 'string'
      THEN jsonb_build_array(statements->'Action')
      WHEN jsonb_typeof(statements->'Action') = 'array'
      THEN statements->'Action'
      END
        AS actions
    FROM
      aws_s3_buckets,
      jsonb_array_elements(
        CASE jsonb_typeof(policy->'Statement')
        WHEN 'string' THEN jsonb_build_array(policy->>'Statement')
        WHEN 'array' THEN policy->'Statement'
        END
      )
        AS statements
    WHERE
      statements->'Effect' = '"Allow"'
  )
    AS flatten_statements,
  jsonb_array_elements(to_jsonb(actions)) AS a,
  jsonb_array_elements(to_jsonb(principals)) AS p
WHERE
  (
    p.value::STRING NOT LIKE '"arn:aws:iam::' || account_id || ':%"'
    OR p.value::STRING = '"*"'
  )
  AND (
      a.value::STRING LIKE '"s3:%*"'
      OR a.value::STRING LIKE '"s3:DeleteObject"'
    );
```

### S3 buckets with replication rules should be enabled

```sql
SELECT
  'S3 buckets with replication rules should be enabled' AS title,
  aws_s3_buckets.account_id,
  aws_s3_buckets.arn AS resource_id,
  CASE
  WHEN r->>'Status' IS DISTINCT FROM 'Enabled' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_s3_buckets,
  jsonb_array_elements(
    CASE jsonb_typeof(replication_rules)
    WHEN 'array' THEN replication_rules
    ELSE '[]'
    END
  )
    AS r;
```

### S3 buckets should have server-side encryption enabled

```sql
SELECT
  'S3 buckets should have server-side encryption enabled' AS title,
  aws_s3_buckets.account_id,
  arn AS resource_id,
  CASE
  WHEN aws_s3_bucket_encryption_rules.bucket_arn IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_s3_buckets
  LEFT JOIN aws_s3_bucket_encryption_rules ON
      aws_s3_bucket_encryption_rules.bucket_arn = aws_s3_buckets.arn;
```


