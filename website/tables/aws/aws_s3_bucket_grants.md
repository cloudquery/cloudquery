# Table: aws_s3_bucket_grants

This table shows data for S3 Bucket Grants.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html

The composite primary key for this table is (**bucket_arn**, **grantee_type**, **grantee_id**, **permission**).

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn (PK)|`utf8`|
|grantee_type (PK)|`utf8`|
|grantee_id (PK)|`utf8`|
|permission (PK)|`utf8`|
|grantee|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


