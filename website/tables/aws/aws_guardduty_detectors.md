# Table: aws_guardduty_detectors

This table shows data for Amazon GuardDuty Detectors.

https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetDetector.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Relations

The following tables depend on aws_guardduty_detectors:
  - [aws_guardduty_detector_filters](aws_guardduty_detector_filters)
  - [aws_guardduty_detector_findings](aws_guardduty_detector_findings)
  - [aws_guardduty_detector_intel_sets](aws_guardduty_detector_intel_sets)
  - [aws_guardduty_detector_ip_sets](aws_guardduty_detector_ip_sets)
  - [aws_guardduty_detector_members](aws_guardduty_detector_members)
  - [aws_guardduty_detector_publishing_destinations](aws_guardduty_detector_publishing_destinations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn|`utf8`|
|id (PK)|`utf8`|
|service_role|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|data_sources|`json`|
|features|`json`|
|finding_publishing_frequency|`utf8`|
|tags|`json`|
|updated_at|`timestamp[us, tz=UTC]`|
|result_metadata|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


