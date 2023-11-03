# Table: aws_accessanalyzer_analyzers

This table shows data for AWS Identity and Access Management (IAM) Access Analyzer Analyzers.

https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_AnalyzerSummary.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_accessanalyzer_analyzers:
  - [aws_accessanalyzer_analyzer_archive_rules](aws_accessanalyzer_analyzer_archive_rules.md)
  - [aws_accessanalyzer_analyzer_findings](aws_accessanalyzer_analyzer_findings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|
|type|`utf8`|
|last_resource_analyzed|`utf8`|
|last_resource_analyzed_at|`timestamp[us, tz=UTC]`|
|status_reason|`json`|
|tags|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure that IAM Access analyzer is enabled for all regions (Automated)

```sql
SELECT
  'Ensure that IAM Access analyzer is enabled for all regions (Automated)'
    AS title,
  ar.account_id,
  ar.region AS resource_id,
  CASE
  WHEN ar.enabled
  AND aregion.region IS NULL
  AND aregion.status IS DISTINCT FROM 'ACTIVE'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_regions AS ar
  LEFT JOIN aws_accessanalyzer_analyzers AS aregion ON
      ar.region = aregion.region;
```


