# Table: aws_rds_clusters

This table shows data for Amazon Relational Database Service (RDS) Clusters.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_rds_clusters:
  - [aws_rds_cluster_backtracks](aws_rds_cluster_backtracks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|activity_stream_kinesis_stream_name|`utf8`|
|activity_stream_kms_key_id|`utf8`|
|activity_stream_mode|`utf8`|
|activity_stream_status|`utf8`|
|allocated_storage|`int64`|
|associated_roles|`json`|
|auto_minor_version_upgrade|`bool`|
|automatic_restart_time|`timestamp[us, tz=UTC]`|
|availability_zones|`list<item: utf8, nullable>`|
|backtrack_consumed_change_records|`int64`|
|backtrack_window|`int64`|
|backup_retention_period|`int64`|
|capacity|`int64`|
|character_set_name|`utf8`|
|clone_group_id|`utf8`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|copy_tags_to_snapshot|`bool`|
|cross_account_clone|`bool`|
|custom_endpoints|`list<item: utf8, nullable>`|
|db_cluster_arn|`utf8`|
|db_cluster_identifier|`utf8`|
|db_cluster_instance_class|`utf8`|
|db_cluster_members|`json`|
|db_cluster_option_group_memberships|`json`|
|db_cluster_parameter_group|`utf8`|
|db_subnet_group|`utf8`|
|db_system_id|`utf8`|
|database_name|`utf8`|
|db_cluster_resource_id|`utf8`|
|deletion_protection|`bool`|
|domain_memberships|`json`|
|earliest_backtrack_time|`timestamp[us, tz=UTC]`|
|earliest_restorable_time|`timestamp[us, tz=UTC]`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`utf8`|
|engine|`utf8`|
|engine_mode|`utf8`|
|engine_version|`utf8`|
|global_write_forwarding_requested|`bool`|
|global_write_forwarding_status|`utf8`|
|hosted_zone_id|`utf8`|
|http_endpoint_enabled|`bool`|
|iam_database_authentication_enabled|`bool`|
|io_optimized_next_allowed_modification_time|`timestamp[us, tz=UTC]`|
|iops|`int64`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|master_user_secret|`json`|
|master_username|`utf8`|
|monitoring_interval|`int64`|
|monitoring_role_arn|`utf8`|
|multi_az|`bool`|
|network_type|`utf8`|
|pending_modified_values|`json`|
|percent_progress|`utf8`|
|performance_insights_enabled|`bool`|
|performance_insights_kms_key_id|`utf8`|
|performance_insights_retention_period|`int64`|
|port|`int64`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|publicly_accessible|`bool`|
|read_replica_identifiers|`list<item: utf8, nullable>`|
|reader_endpoint|`utf8`|
|replication_source_identifier|`utf8`|
|scaling_configuration_info|`json`|
|serverless_v2_scaling_configuration|`json`|
|status|`utf8`|
|storage_encrypted|`bool`|
|storage_type|`utf8`|
|vpc_security_groups|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon Aurora clusters should have backtracking enabled

```sql
SELECT
  'Amazon Aurora clusters should have backtracking enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN backtrack_window IS NULL THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_rds_clusters
WHERE
  engine IN ('aurora', 'aurora-mysql', 'mysql');
```

### IAM authentication should be configured for RDS clusters

```sql
SELECT
  'IAM authentication should be configured for RDS clusters' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN iam_database_authentication_enabled IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_clusters;
```

### RDS clusters should have deletion protection enabled

```sql
SELECT
  'RDS clusters should have deletion protection enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN deletion_protection IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_clusters;
```

### RDS databases and clusters should not use a database engine default port

```sql
(
  SELECT
    'RDS databases and clusters should not use a database engine default port'
      AS title,
    account_id,
    arn AS resource_id,
    CASE
    WHEN (engine IN ('aurora', 'aurora-mysql', 'mysql') AND port = 3306)
    OR (engine LIKE '%postgres%' AND port = 5432)
    THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    aws_rds_clusters
)
UNION
  (
    SELECT
      'RDS databases and clusters should not use a database engine default port'
        AS title,
      account_id,
      arn AS resource_id,
      CASE
      WHEN (
        engine IN ('aurora', 'aurora-mysql', 'mariadb', 'mysql')
        AND db_instance_port = 3306
      )
      OR (engine LIKE '%postgres%' AND db_instance_port = 5432)
      OR (engine LIKE '%oracle%' AND db_instance_port = 1521)
      OR (engine LIKE '%sqlserver%' AND db_instance_port = 1433)
      THEN 'fail'
      ELSE 'pass'
      END
        AS status
    FROM
      aws_rds_instances
  );
```

### RDS DB clusters should be configured for multiple Availability Zones

```sql
SELECT
  'RDS DB clusters should be configured for multiple Availability Zones'
    AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN multi_az IS NOT true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_rds_clusters;
```

### RDS DB clusters should be configured to copy tags to snapshots

```sql
SELECT
  'RDS DB clusters should be configured to copy tags to snapshots' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN copy_tags_to_snapshot IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_clusters;
```

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


