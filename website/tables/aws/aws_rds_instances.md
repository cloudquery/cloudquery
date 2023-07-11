# Table: aws_rds_instances

This table shows data for Amazon Relational Database Service (RDS) Instances.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|processor_features|`json`|
|tags|`json`|
|activity_stream_engine_native_audit_fields_included|`bool`|
|activity_stream_kinesis_stream_name|`utf8`|
|activity_stream_kms_key_id|`utf8`|
|activity_stream_mode|`utf8`|
|activity_stream_policy_status|`utf8`|
|activity_stream_status|`utf8`|
|allocated_storage|`int64`|
|associated_roles|`json`|
|auto_minor_version_upgrade|`bool`|
|automatic_restart_time|`timestamp[us, tz=UTC]`|
|automation_mode|`utf8`|
|availability_zone|`utf8`|
|aws_backup_recovery_point_arn|`utf8`|
|backup_retention_period|`int64`|
|backup_target|`utf8`|
|ca_certificate_identifier|`utf8`|
|certificate_details|`json`|
|character_set_name|`utf8`|
|copy_tags_to_snapshot|`bool`|
|custom_iam_instance_profile|`utf8`|
|customer_owned_ip_enabled|`bool`|
|db_cluster_identifier|`utf8`|
|db_instance_arn|`utf8`|
|db_instance_automated_backups_replications|`json`|
|db_instance_class|`utf8`|
|db_instance_identifier|`utf8`|
|db_instance_status|`utf8`|
|db_name|`utf8`|
|db_parameter_groups|`json`|
|db_security_groups|`json`|
|db_subnet_group|`json`|
|db_system_id|`utf8`|
|db_instance_port|`int64`|
|dbi_resource_id|`utf8`|
|deletion_protection|`bool`|
|domain_memberships|`json`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`json`|
|engine|`utf8`|
|engine_version|`utf8`|
|enhanced_monitoring_resource_arn|`utf8`|
|iam_database_authentication_enabled|`bool`|
|instance_create_time|`timestamp[us, tz=UTC]`|
|iops|`int64`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|license_model|`utf8`|
|listener_endpoint|`json`|
|master_user_secret|`json`|
|master_username|`utf8`|
|max_allocated_storage|`int64`|
|monitoring_interval|`int64`|
|monitoring_role_arn|`utf8`|
|multi_az|`bool`|
|nchar_character_set_name|`utf8`|
|network_type|`utf8`|
|option_group_memberships|`json`|
|pending_modified_values|`json`|
|performance_insights_enabled|`bool`|
|performance_insights_kms_key_id|`utf8`|
|performance_insights_retention_period|`int64`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|promotion_tier|`int64`|
|publicly_accessible|`bool`|
|read_replica_db_cluster_identifiers|`list<item: utf8, nullable>`|
|read_replica_db_instance_identifiers|`list<item: utf8, nullable>`|
|read_replica_source_db_cluster_identifier|`utf8`|
|read_replica_source_db_instance_identifier|`utf8`|
|replica_mode|`utf8`|
|resume_full_automation_mode_time|`timestamp[us, tz=UTC]`|
|secondary_availability_zone|`utf8`|
|status_infos|`json`|
|storage_encrypted|`bool`|
|storage_throughput|`int64`|
|storage_type|`utf8`|
|tde_credential_arn|`utf8`|
|timezone|`utf8`|
|vpc_security_groups|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Database logging should be enabled

```sql
SELECT
  'Database logging should be enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN enabled_cloudwatch_logs_exports IS NULL
  OR (
      engine IN ('aurora', 'aurora-mysql', 'mariadb', 'mysql')
      AND NOT
          (enabled_cloudwatch_logs_exports @> '{audit,error,general,slowquery}')
    )
  OR (
      engine LIKE '%postgres%'
      AND NOT (enabled_cloudwatch_logs_exports @> '{postgresql,upgrade}')
    )
  OR (
      engine LIKE '%oracle%'
      AND NOT
          (enabled_cloudwatch_logs_exports @> '{alert,audit,trace,listener}')
    )
  OR (
      engine LIKE '%sqlserver%'
      AND NOT (enabled_cloudwatch_logs_exports @> '{error,agent}')
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```

### Enhanced monitoring should be configured for RDS DB instances and clusters

```sql
SELECT
  'Enhanced monitoring should be configured for RDS DB instances and clusters'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN enhanced_monitoring_resource_arn IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```

### IAM authentication should be configured for RDS instances

```sql
SELECT
  'IAM authentication should be configured for RDS instances' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN iam_database_authentication_enabled IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```

### RDS automatic minor version upgrades should be enabled

```sql
SELECT
  'RDS automatic minor version upgrades should be enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN auto_minor_version_upgrade IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
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

### RDS DB instances should be configured to copy tags to snapshots

```sql
SELECT
  'RDS DB instances should be configured to copy tags to snapshots' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN copy_tags_to_snapshot IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```

### RDS DB instances should be configured with multiple Availability Zones

```sql
SELECT
  'RDS DB instances should be configured with multiple Availability Zones'
    AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN multi_az IS NOT true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_rds_instances;
```

### RDS DB instances should have deletion protection enabled

```sql
SELECT
  'RDS DB instances should have deletion protection enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN deletion_protection IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```

### RDS DB instances should have encryption at rest enabled

```sql
SELECT
  'RDS DB instances should have encryption at rest enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN storage_encrypted IS NOT true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_rds_instances;
```

### RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration

```sql
SELECT
  'RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration'
    AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN publicly_accessible IS true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_rds_instances;
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

### RDS instances should be deployed in a VPC

```sql
SELECT
  'RDS instances should be deployed in a VPC' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (db_subnet_group->>'VpcId') IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_instances;
```


