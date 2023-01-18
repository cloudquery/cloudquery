# Table: aws_docdb_clusters

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_docdb_clusters:
  - [aws_docdb_cluster_snapshots](aws_docdb_cluster_snapshots.md)
  - [aws_docdb_instances](aws_docdb_instances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|associated_roles|JSON|
|availability_zones|StringArray|
|backup_retention_period|Int|
|clone_group_id|String|
|cluster_create_time|Timestamp|
|db_cluster_arn|String|
|db_cluster_identifier|String|
|db_cluster_members|JSON|
|db_cluster_parameter_group|String|
|db_subnet_group|String|
|db_cluster_resource_id|String|
|deletion_protection|Bool|
|earliest_restorable_time|Timestamp|
|enabled_cloudwatch_logs_exports|StringArray|
|endpoint|String|
|engine|String|
|engine_version|String|
|hosted_zone_id|String|
|kms_key_id|String|
|latest_restorable_time|Timestamp|
|master_username|String|
|multi_az|Bool|
|percent_progress|String|
|port|Int|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|read_replica_identifiers|StringArray|
|reader_endpoint|String|
|replication_source_identifier|String|
|status|String|
|storage_encrypted|Bool|
|vpc_security_groups|JSON|