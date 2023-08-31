# Table: aws_cloudfront_distributions

This table shows data for Cloudfront Distributions.

https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|distribution_config|`json`|
|domain_name|`utf8`|
|id|`utf8`|
|in_progress_invalidation_batches|`int64`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|active_trusted_key_groups|`json`|
|active_trusted_signers|`json`|
|alias_icp_recordals|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### CloudFront distributions should have logging enabled

```sql
SELECT
  'CloudFront distributions should have logging enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (distribution_config->'Logging'->>'Enabled')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudfront_distributions;
```

### Find all CloudFront distributions

```sql
SELECT
  'Find all CloudFront distributions' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_cloudfront_distributions;
```

### API Gateway should be associated with an AWS WAF web ACL

```sql
SELECT
  'API Gateway should be associated with an AWS WAF web ACL' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN distribution_config->>'WebACLId' = '' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudfront_distributions;
```

### CloudFront distributions should have a default root object configured

```sql
SELECT
  'CloudFront distributions should have a default root object configured'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN distribution_config->>'DefaultRootObject' = '' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudfront_distributions;
```

### Disabled CloudFront distribution

```sql
SELECT
  'Disabled CloudFront distribution' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_cloudfront_distributions
WHERE
  (distribution_config->>'Enabled')::BOOL IS NOT true;
```

### CloudFront distributions should have origin access identity enabled

```sql
SELECT
  'CloudFront distributions should have origin access identity enabled'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN o->>'DomainName' LIKE '%s3.amazonaws.com'
  AND o->'S3OriginConfig'->>'OriginAccessIdentity' = ''
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudfront_distributions,
  jsonb_array_elements(distribution_config->'Origins'->'Items') AS o;
```

### CloudFront distributions should have origin failover configured

```sql
WITH
  origin_groups
    AS (
      SELECT
        acd.arn, distribution_config->'OriginGroups'->'Items' AS ogs
      FROM
        aws_cloudfront_distributions AS acd
    ),
  oids
    AS (
      SELECT
        DISTINCT
        account_id,
        acd.arn AS resource_id,
        CASE
        WHEN o.ogs = 'null'
        OR o.ogs->'Members'->'Items' = 'null'
        OR jsonb_array_length(o.ogs->'Members'->'Items') = 0
        THEN 'fail'
        ELSE 'pass'
        END
          AS status
      FROM
        aws_cloudfront_distributions AS acd
        LEFT JOIN origin_groups AS o ON o.arn = acd.arn
    )
SELECT
  'CloudFront distributions should have origin failover configured' AS title,
  account_id,
  resource_id,
  status
FROM
  oids;
```

### CloudFront distributions should require encryption in transit

```sql
WITH
  cachebeviors
    AS (
      SELECT
        DISTINCT arn, account_id
      FROM
        (
          SELECT
            arn, account_id, d AS cachebehavior
          FROM
            aws_cloudfront_distributions,
            jsonb_array_elements(distribution_config->'CacheBehaviors'->'Items')
              AS d
          WHERE
            distribution_config->'CacheBehaviors'->'Items' != 'null'
          UNION
            SELECT
              arn,
              account_id,
              distribution_config->'DefaultCacheBehavior' AS cachebehavior
            FROM
              aws_cloudfront_distributions
        )
          AS cachebeviors
      WHERE
        cachebehavior->>'ViewerProtocolPolicy' = 'allow-all'
    )
SELECT
  'CloudFront distributions should require encryption in transit' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  cachebeviors;
```


