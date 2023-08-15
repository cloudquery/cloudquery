# Table: aws_rds_db_parameter_groups

This table shows data for Amazon Relational Database Service (RDS) DB Parameter Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_rds_db_parameter_groups:
  - [aws_rds_db_parameter_group_db_parameters](aws_rds_db_parameter_group_db_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|db_parameter_group_arn|`utf8`|
|db_parameter_group_family|`utf8`|
|db_parameter_group_name|`utf8`|
|description|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### An RDS event notifications subscription should be configured for critical database parameter group events

```sql
WITH
  any_category
    AS (
      SELECT
        DISTINCT true AS any_category
      FROM
        aws_rds_event_subscriptions
      WHERE
        (source_type IS NULL OR source_type = 'db-parameter-group')
        AND event_categories_list IS NULL
    ),
  any_source_id
    AS (
      SELECT
        COALESCE(array_agg(category), '{}'::STRING[]) AS any_source_categories
      FROM
        aws_rds_event_subscriptions, unnest(event_categories_list) AS category
      WHERE
        source_type = 'db-parameter-group' AND event_categories_list IS NOT NULL
    ),
  specific_categories
    AS (
      SELECT
        source_id, array_agg(category) AS specific_cats
      FROM
        aws_rds_event_subscriptions,
        unnest(source_ids_list) AS source_id,
        unnest(event_categories_list) AS category
      WHERE
        source_type = 'db-parameter-group'
      GROUP BY
        source_id
    )
SELECT
  'An RDS event notifications subscription should be configured for critical database parameter group events'
    AS title,
  aws_rds_db_parameter_groups.account_id,
  aws_rds_db_parameter_groups.arn AS resource_id,
  CASE
  WHEN any_category IS NOT true
  AND NOT (any_source_categories @> '{"configuration change"}')
  AND (
      specific_cats IS NULL
      OR NOT (specific_cats @> '{"configuration change"}')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_db_parameter_groups
  LEFT JOIN any_category ON true
  INNER JOIN any_source_id ON true
  LEFT JOIN specific_categories ON
      aws_rds_db_parameter_groups.db_parameter_group_name
      = specific_categories.source_id;
```


