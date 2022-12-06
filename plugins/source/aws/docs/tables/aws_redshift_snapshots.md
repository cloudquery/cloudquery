# Table: aws_redshift_snapshots

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Snapshot.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_redshift_clusters](aws_redshift_clusters.md).


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
|tags|JSON|
|accounts_with_restore_access|JSON|
|actual_incremental_backup_size_in_mega_bytes|Float|
|availability_zone|String|
|backup_progress_in_mega_bytes|Float|
|cluster_create_time|Timestamp|
|cluster_identifier|String|
|cluster_version|String|
|current_backup_rate_in_mega_bytes_per_second|Float|
|db_name|String|
|elapsed_time_in_seconds|Int|
|encrypted|Bool|
|encrypted_with_hsm|Bool|
|engine_full_version|String|
|enhanced_vpc_routing|Bool|
|estimated_seconds_to_completion|Int|
|kms_key_id|String|
|maintenance_track_name|String|
|manual_snapshot_remaining_days|Int|
|manual_snapshot_retention_period|Int|
|master_username|String|
|node_type|String|
|number_of_nodes|Int|
|owner_account|String|
|port|Int|
|restorable_node_types|StringArray|
|snapshot_create_time|Timestamp|
|snapshot_identifier|String|
|snapshot_retention_start_time|Timestamp|
|snapshot_type|String|
|source_region|String|
|status|String|
|total_backup_size_in_mega_bytes|Float|
|vpc_id|String|