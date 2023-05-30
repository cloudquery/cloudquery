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
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
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