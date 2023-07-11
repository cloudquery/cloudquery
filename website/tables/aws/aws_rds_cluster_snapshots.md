# Table: aws_rds_cluster_snapshots

This table shows data for Amazon Relational Database Service (RDS) Cluster Snapshots.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html

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
|attributes|`json`|
|allocated_storage|`int64`|
|availability_zones|`list<item: utf8, nullable>`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|db_cluster_identifier|`utf8`|
|db_cluster_snapshot_arn|`utf8`|
|db_cluster_snapshot_identifier|`utf8`|
|db_system_id|`utf8`|
|engine|`utf8`|
|engine_mode|`utf8`|
|engine_version|`utf8`|
|iam_database_authentication_enabled|`bool`|
|kms_key_id|`utf8`|
|license_model|`utf8`|
|master_username|`utf8`|
|percent_progress|`int64`|
|port|`int64`|
|snapshot_create_time|`timestamp[us, tz=UTC]`|
|snapshot_type|`utf8`|
|source_db_cluster_snapshot_arn|`utf8`|
|status|`utf8`|
|storage_encrypted|`bool`|
|storage_type|`utf8`|
|vpc_id|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### RDS cluster snapshots and database snapshots should be encrypted at rest

```sql
(
  SELECT
    'RDS cluster snapshots and database snapshots should be encrypted at rest'
      AS title,
    account_id,
    arn AS resource_id,
    CASE
    WHEN storage_encrypted IS NOT true THEN 'fail'
    ELSE 'pass'
    END
      AS status
  FROM
    aws_rds_cluster_snapshots
)
UNION
  (
    SELECT
      'RDS cluster snapshots and database snapshots should be encrypted at rest'
        AS title,
      account_id,
      arn AS resource_id,
      CASE WHEN encrypted IS NOT true THEN 'fail' ELSE 'pass' END AS status
    FROM
      aws_rds_db_snapshots
  );
```

### RDS snapshots should be private

```sql
SELECT
  'RDS snapshots should be private' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN attrs->>'AttributeName' IS NOT DISTINCT FROM 'restore'
  AND (attrs->'AttributeValues')::JSONB ? 'all'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_rds_cluster_snapshots, jsonb_array_elements(attributes) AS attrs;
```


