# Table: aws_securityhub_hubs

This table shows data for AWS Security Hub Hubs.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html

The composite primary key for this table is (**account_id**, **region**, **hub_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|auto_enable_controls|`bool`|
|control_finding_generator|`utf8`|
|hub_arn (PK)|`utf8`|
|subscribed_at|`timestamp[us, tz=UTC]`|
|result_metadata|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

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


