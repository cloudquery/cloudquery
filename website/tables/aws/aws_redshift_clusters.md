# Table: aws_redshift_clusters

This table shows data for Redshift Clusters.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_redshift_clusters:
  - [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups)
  - [aws_redshift_endpoint_access](aws_redshift_endpoint_access)
  - [aws_redshift_endpoint_authorization](aws_redshift_endpoint_authorization)
  - [aws_redshift_snapshots](aws_redshift_snapshots)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|logging_status|`json`|
|tags|`json`|
|allow_version_upgrade|`bool`|
|aqua_configuration|`json`|
|automated_snapshot_retention_period|`int64`|
|availability_zone|`utf8`|
|availability_zone_relocation_status|`utf8`|
|cluster_availability_status|`utf8`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|cluster_identifier|`utf8`|
|cluster_namespace_arn|`utf8`|
|cluster_nodes|`json`|
|cluster_public_key|`utf8`|
|cluster_revision_number|`utf8`|
|cluster_security_groups|`json`|
|cluster_snapshot_copy_status|`json`|
|cluster_status|`utf8`|
|cluster_subnet_group_name|`utf8`|
|cluster_version|`utf8`|
|custom_domain_certificate_arn|`utf8`|
|custom_domain_certificate_expiry_date|`timestamp[us, tz=UTC]`|
|custom_domain_name|`utf8`|
|db_name|`utf8`|
|data_transfer_progress|`json`|
|default_iam_role_arn|`utf8`|
|deferred_maintenance_windows|`json`|
|elastic_ip_status|`json`|
|elastic_resize_number_of_node_options|`utf8`|
|encrypted|`bool`|
|endpoint|`json`|
|enhanced_vpc_routing|`bool`|
|expected_next_snapshot_schedule_time|`timestamp[us, tz=UTC]`|
|expected_next_snapshot_schedule_time_status|`utf8`|
|hsm_status|`json`|
|iam_roles|`json`|
|kms_key_id|`utf8`|
|maintenance_track_name|`utf8`|
|manual_snapshot_retention_period|`int64`|
|master_username|`utf8`|
|modify_status|`utf8`|
|next_maintenance_window_start_time|`timestamp[us, tz=UTC]`|
|node_type|`utf8`|
|number_of_nodes|`int64`|
|pending_actions|`list<item: utf8, nullable>`|
|pending_modified_values|`json`|
|preferred_maintenance_window|`utf8`|
|publicly_accessible|`bool`|
|reserved_node_exchange_status|`json`|
|resize_info|`json`|
|restore_status|`json`|
|snapshot_schedule_identifier|`utf8`|
|snapshot_schedule_state|`utf8`|
|total_storage_capacity_in_mega_bytes|`int64`|
|vpc_id|`utf8`|
|vpc_security_groups|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Amazon Redshift clusters should prohibit public access

```sql
SELECT
  'Amazon Redshift clusters should prohibit public access' AS title,
  account_id,
  arn AS resource_id,
  CASE WHEN publicly_accessible IS true THEN 'fail' ELSE 'pass' END AS status
FROM
  aws_redshift_clusters;
```

### Connections to Amazon Redshift clusters should be encrypted in transit

```sql
SELECT
  'Connections to Amazon Redshift clusters should be encrypted in transit'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_redshift_clusters AS rsc
WHERE
  EXISTS(
    SELECT
      1
    FROM
      aws_redshift_cluster_parameter_groups AS rscpg
      INNER JOIN aws_redshift_cluster_parameters AS rscp ON
          rscpg.cluster_arn = rscp.cluster_arn
    WHERE
      rsc.arn = rscpg.cluster_arn
      AND (
          rscp.parameter_name = 'require_ssl'
          AND rscp.parameter_value = 'false'
        )
      OR (rscp.parameter_name = 'require_ssl' AND rscp.parameter_value IS NULL)
      OR NOT
          EXISTS(
            (
              SELECT
                1
              FROM
                aws_redshift_cluster_parameters
              WHERE
                cluster_arn = rscpg.cluster_arn
                AND parameter_name = 'require_ssl'
            )
          )
  );
```

### Amazon Redshift clusters should have audit logging enabled

```sql
SELECT
  'Amazon Redshift clusters should have audit logging enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN jsonb_typeof(logging_status->'LoggingEnabled') IS NULL
  OR (
      jsonb_typeof(logging_status->'LoggingEnabled') IS NOT NULL
      AND (logging_status->>'LoggingEnabled')::BOOL IS false
    )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_redshift_clusters;
```

### Amazon Redshift clusters should have automatic snapshots enabled

```sql
SELECT
  'Amazon Redshift clusters should have automatic snapshots enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN automated_snapshot_retention_period < 7
  OR automated_snapshot_retention_period IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_redshift_clusters;
```

### Amazon Redshift should have automatic upgrades to major versions enabled

```sql
SELECT
  'Amazon Redshift should have automatic upgrades to major versions enabled'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN allow_version_upgrade IS false OR allow_version_upgrade IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_redshift_clusters;
```

### Amazon Redshift clusters should use enhanced VPC routing

```sql
SELECT
  'Amazon Redshift clusters should use enhanced VPC routing' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN enhanced_vpc_routing IS false OR enhanced_vpc_routing IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_redshift_clusters;
```


