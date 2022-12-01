# Table: aws_redshift_clusters

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_redshift_clusters:
  - [aws_redshift_snapshots](aws_redshift_snapshots.md)
  - [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|logging_status|JSON|
|allow_version_upgrade|Bool|
|aqua_configuration|JSON|
|automated_snapshot_retention_period|Int|
|availability_zone|String|
|availability_zone_relocation_status|String|
|cluster_availability_status|String|
|cluster_create_time|Timestamp|
|cluster_identifier|String|
|cluster_namespace_arn|String|
|cluster_nodes|JSON|
|cluster_public_key|String|
|cluster_revision_number|String|
|cluster_security_groups|JSON|
|cluster_snapshot_copy_status|JSON|
|cluster_status|String|
|cluster_subnet_group_name|String|
|cluster_version|String|
|db_name|String|
|data_transfer_progress|JSON|
|default_iam_role_arn|String|
|deferred_maintenance_windows|JSON|
|elastic_ip_status|JSON|
|elastic_resize_number_of_node_options|String|
|encrypted|Bool|
|endpoint|JSON|
|enhanced_vpc_routing|Bool|
|expected_next_snapshot_schedule_time|Timestamp|
|expected_next_snapshot_schedule_time_status|String|
|hsm_status|JSON|
|iam_roles|JSON|
|kms_key_id|String|
|maintenance_track_name|String|
|manual_snapshot_retention_period|Int|
|master_username|String|
|modify_status|String|
|next_maintenance_window_start_time|Timestamp|
|node_type|String|
|number_of_nodes|Int|
|pending_actions|StringArray|
|pending_modified_values|JSON|
|preferred_maintenance_window|String|
|publicly_accessible|Bool|
|reserved_node_exchange_status|JSON|
|resize_info|JSON|
|restore_status|JSON|
|snapshot_schedule_identifier|String|
|snapshot_schedule_state|String|
|tags|JSON|
|total_storage_capacity_in_mega_bytes|Int|
|vpc_id|String|
|vpc_security_groups|JSON|