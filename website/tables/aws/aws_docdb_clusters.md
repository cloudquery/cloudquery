# Table: aws_docdb_clusters

This table shows data for Amazon DocumentDB Clusters.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_docdb_clusters:
  - [aws_docdb_cluster_snapshots](aws_docdb_cluster_snapshots)
  - [aws_docdb_instances](aws_docdb_instances)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|associated_roles|`json`|
|availability_zones|`list<item: utf8, nullable>`|
|backup_retention_period|`int64`|
|clone_group_id|`utf8`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|db_cluster_arn|`utf8`|
|db_cluster_identifier|`utf8`|
|db_cluster_members|`json`|
|db_cluster_parameter_group|`utf8`|
|db_subnet_group|`utf8`|
|db_cluster_resource_id|`utf8`|
|deletion_protection|`bool`|
|earliest_restorable_time|`timestamp[us, tz=UTC]`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|hosted_zone_id|`utf8`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|master_username|`utf8`|
|multi_az|`bool`|
|percent_progress|`utf8`|
|port|`int64`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|read_replica_identifiers|`list<item: utf8, nullable>`|
|reader_endpoint|`utf8`|
|replication_source_identifier|`utf8`|
|status|`utf8`|
|storage_encrypted|`bool`|
|vpc_security_groups|`json`|