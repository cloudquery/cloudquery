# Table: aws_redshift_snapshots

This table shows data for Redshift Snapshots.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Snapshot.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|accounts_with_restore_access|`json`|
|actual_incremental_backup_size_in_mega_bytes|`float64`|
|availability_zone|`utf8`|
|backup_progress_in_mega_bytes|`float64`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|cluster_identifier|`utf8`|
|cluster_version|`utf8`|
|current_backup_rate_in_mega_bytes_per_second|`float64`|
|db_name|`utf8`|
|elapsed_time_in_seconds|`int64`|
|encrypted|`bool`|
|encrypted_with_hsm|`bool`|
|engine_full_version|`utf8`|
|enhanced_vpc_routing|`bool`|
|estimated_seconds_to_completion|`int64`|
|kms_key_id|`utf8`|
|maintenance_track_name|`utf8`|
|manual_snapshot_remaining_days|`int64`|
|manual_snapshot_retention_period|`int64`|
|master_username|`utf8`|
|node_type|`utf8`|
|number_of_nodes|`int64`|
|owner_account|`utf8`|
|port|`int64`|
|restorable_node_types|`list<item: utf8, nullable>`|
|snapshot_create_time|`timestamp[us, tz=UTC]`|
|snapshot_identifier|`utf8`|
|snapshot_retention_start_time|`timestamp[us, tz=UTC]`|
|snapshot_type|`utf8`|
|source_region|`utf8`|
|status|`utf8`|
|total_backup_size_in_mega_bytes|`float64`|
|vpc_id|`utf8`|