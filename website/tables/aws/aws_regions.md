# Table: aws_regions

This table shows data for Regions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|enabled|`bool`|
|partition|`utf8`|
|region|`utf8`|
|endpoint|`utf8`|
|opt_in_status|`utf8`|
|region_name|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that IAM Access analyzer is enabled for all regions (Automated)

```sql
WITH
  regions_with_enabled_accessanalyzer
    AS (
      SELECT
        ar.region AS analyzed_region
      FROM
        aws_regions AS ar
        LEFT JOIN aws_accessanalyzer_analyzers AS aaaa ON
            ar.region = aaaa.region
      WHERE
        aaaa.status = 'ACTIVE'
    )
SELECT
  'Ensure that IAM Access analyzer is enabled for all regions (Automated)'
    AS title,
  account_id,
  region AS resource_id,
  CASE
  WHEN aregion.analyzed_region IS NULL AND ar.enabled = true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_regions AS ar
  LEFT JOIN regions_with_enabled_accessanalyzer AS aregion ON
      ar.region = aregion.analyzed_region;
```

### GuardDuty should be enabled

```sql
WITH
  enabled_detector_regions
    AS (
      SELECT
        account_id, region
      FROM
        aws_guardduty_detectors
      WHERE
        status = 'ENABLED'
    )
SELECT
  'GuardDuty should be enabled' AS title,
  r.account_id,
  r.region AS resource_id,
  CASE
  WHEN enabled = true AND e.region IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_regions AS r
  LEFT JOIN enabled_detector_regions AS e ON
      e.region = r.region AND e.account_id = r.account_id
UNION
  SELECT
    'GuardDuty should be enabled (detectors)' AS title,
    account_id,
    region AS resource_id,
    CASE
    WHEN data_sources->'S3Logs'->>'Status' != 'ENABLED'
    AND data_sources->'DNSLogs'->>'Status' != 'ENABLED'
    AND data_sources->'CloudTrail'->>'Status' != 'ENABLED'
    AND data_sources->'FlowLogs'->>'Status' != 'ENABLED'
    THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    aws_guardduty_detectors
  WHERE
    status = 'ENABLED';
```

### SecurityHub should be enabled

```sql
WITH
  enabled_securityhub_regions
    AS (SELECT account_id, region FROM aws_securityhub_hubs)
SELECT
  'SecurityHub should be enabled' AS title,
  r.account_id,
  r.region AS resource_id,
  CASE
  WHEN r.enabled = true AND e.region IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_regions AS r
  LEFT JOIN enabled_securityhub_regions AS e ON
      e.region = r.region AND e.account_id = r.account_id;
```


