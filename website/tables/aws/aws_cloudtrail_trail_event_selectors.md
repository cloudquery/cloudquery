# Table: aws_cloudtrail_trail_event_selectors

This table shows data for AWS CloudTrail Trail Event Selectors.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_EventSelector.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudtrail_trails](aws_cloudtrail_trails).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|trail_arn|`utf8`|
|data_resources|`json`|
|exclude_management_event_sources|`list<item: utf8, nullable>`|
|include_management_events|`bool`|
|read_write_type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure CloudTrail is enabled in all regions

```sql
SELECT
  'Ensure CloudTrail is enabled in all regions' AS title,
  aws_cloudtrail_trails.account_id,
  arn AS resource_id,
  CASE
  WHEN is_multi_region_trail = false
  OR (
      is_multi_region_trail = true
      AND (read_write_type != 'All' OR include_management_events = false)
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_cloudtrail_trails
  INNER JOIN aws_cloudtrail_trail_event_selectors ON
      aws_cloudtrail_trails.arn = aws_cloudtrail_trail_event_selectors.trail_arn
      AND aws_cloudtrail_trails.region
        = aws_cloudtrail_trail_event_selectors.region
      AND aws_cloudtrail_trails.account_id
        = aws_cloudtrail_trail_event_selectors.account_id;
```


