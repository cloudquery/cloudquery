# Table: aws_rds_event_subscriptions

This table shows data for Amazon Relational Database Service (RDS) Event Subscriptions.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_EventSubscription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|cust_subscription_id|`utf8`|
|customer_aws_id|`utf8`|
|enabled|`bool`|
|event_categories_list|`list<item: utf8, nullable>`|
|event_subscription_arn|`utf8`|
|sns_topic_arn|`utf8`|
|source_ids_list|`list<item: utf8, nullable>`|
|source_type|`utf8`|
|status|`utf8`|
|subscription_creation_time|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### An RDS event notifications subscription should be configured for critical cluster events

```sql
WITH
  any_category
    AS (
      SELECT
        DISTINCT true AS any_category
      FROM
        aws_rds_event_subscriptions
      WHERE
        (source_type IS NULL OR source_type = 'db-cluster')
        AND event_categories_list IS NULL
    ),
  any_source_id
    AS (
      SELECT
        COALESCE(array_agg(category), '{}'::STRING[]) AS any_source_categories
      FROM
        aws_rds_event_subscriptions, unnest(event_categories_list) AS category
      WHERE
        source_type = 'db-cluster' AND event_categories_list IS NOT NULL
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
        source_type = 'db-cluster'
      GROUP BY
        source_id
    )
SELECT
  'An RDS event notifications subscription should be configured for critical cluster events'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN any_category IS NOT true
  AND NOT (any_source_categories @> '{"failure","maintenance"}')
  AND (
      specific_cats IS NULL
      OR NOT (specific_cats @> '{"failure","maintenance"}')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_clusters
  LEFT JOIN any_category ON true
  INNER JOIN any_source_id ON true
  LEFT JOIN specific_categories ON
      db_cluster_identifier = specific_categories.source_id;
```

### An RDS event notifications subscription should be configured for critical database instance events

```sql
WITH
  any_category
    AS (
      SELECT
        DISTINCT true AS any_category
      FROM
        aws_rds_event_subscriptions
      WHERE
        (source_type IS NULL OR source_type = 'db-instance')
        AND event_categories_list IS NULL
    ),
  any_source_id
    AS (
      SELECT
        COALESCE(array_agg(category), '{}'::STRING[]) AS any_source_categories
      FROM
        aws_rds_event_subscriptions, unnest(event_categories_list) AS category
      WHERE
        source_type = 'db-instance' AND event_categories_list IS NOT NULL
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
        source_type = 'db-instance'
      GROUP BY
        source_id
    )
SELECT
  'An RDS event notifications subscription should be configured for critical database instance events'
    AS title,
  aws_rds_instances.account_id,
  aws_rds_instances.arn AS resource_id,
  CASE
  WHEN any_category IS NOT true
  AND NOT
      (
        any_source_categories
        @> '{"maintenance","configuration change","failure"}'
      )
  AND (
      specific_cats IS NULL
      OR NOT
          (specific_cats @> '{"maintenance","configuration change","failure"}')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances
  LEFT JOIN any_category ON true
  INNER JOIN any_source_id ON true
  LEFT JOIN specific_categories ON
      db_instance_identifier = specific_categories.source_id;
```

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

### An RDS event notifications subscription should be configured for critical database security group events

```sql
WITH
  any_category
    AS (
      SELECT
        DISTINCT true AS any_category
      FROM
        aws_rds_event_subscriptions
      WHERE
        (source_type IS NULL OR source_type = 'db-security-group')
        AND event_categories_list IS NULL
    ),
  any_source_id
    AS (
      SELECT
        COALESCE(array_agg(category), '{}'::STRING[]) AS any_source_categories
      FROM
        aws_rds_event_subscriptions, unnest(event_categories_list) AS category
      WHERE
        source_type = 'db-security-group' AND event_categories_list IS NOT NULL
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
        source_type = 'db-security-group'
      GROUP BY
        source_id
    )
SELECT
  'An RDS event notifications subscription should be configured for critical database security group events'
    AS title,
  aws_rds_db_security_groups.account_id,
  aws_rds_db_security_groups.arn AS resource_id,
  CASE
  WHEN any_category IS NOT true
  AND NOT (any_source_categories @> '{"configuration change","failure"}')
  AND (
      specific_cats IS NULL
      OR NOT (specific_cats @> '{"configuration change","failure"}')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_db_security_groups
  LEFT JOIN any_category ON true
  INNER JOIN any_source_id ON true
  LEFT JOIN specific_categories ON
      aws_rds_db_security_groups.db_security_group_name
      = specific_categories.source_id;
```


