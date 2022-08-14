
# Table: aws_redshift_snapshots
Describes a snapshot.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_redshift_clusters table (FK)|
|arn|text|ARN of the snapshot.|
|actual_incremental_backup_size|float|The size of the incremental backup in megabytes.|
|availability_zone|text|The Availability Zone in which the cluster was created.|
|backup_progress|float|The number of megabytes that have been transferred to the snapshot backup.|
|cluster_create_time|timestamp without time zone|The time (UTC) when the cluster was originally created.|
|cluster_identifier|text|The identifier of the cluster for which the snapshot was taken.|
|cluster_version|text|The version ID of the Amazon Redshift engine that is running on the cluster.|
|current_backup_rate|float|The number of megabytes per second being transferred to the snapshot backup. Returns 0 for a completed backup.|
|db_name|text|The name of the database that was created when the cluster was created.|
|elapsed_time|bigint|The amount of time an in-progress snapshot backup has been running, or the amount of time it took a completed backup to finish, in seconds.|
|encrypted|boolean|If true, the data in the snapshot is encrypted at rest.|
|encrypted_with_hsm|boolean|A boolean that indicates whether the snapshot data is encrypted using the HSM keys of the source cluster.|
|engine_full_version|text|The cluster version of the cluster used to create the snapshot.|
|enhanced_vpc_routing|boolean|An option that specifies whether to create the cluster with enhanced VPC routing enabled.|
|estimated_seconds_to_completion|bigint|The estimate of the time remaining before the snapshot backup will complete. Returns 0 for a completed backup.|
|kms_key_id|text|The AWS Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.|
|maintenance_track_name|text|The name of the maintenance track for the snapshot.|
|manual_snapshot_remaining_days|integer|The number of days until a manual snapshot will pass its retention period.|
|manual_snapshot_retention_period|integer|The number of days that a manual snapshot is retained|
|master_username|text|The master user name for the cluster.|
|node_type|text|The node type of the nodes in the cluster.|
|number_of_nodes|integer|The number of nodes in the cluster.|
|owner_account|text|For manual snapshots, the AWS customer account used to create or copy the snapshot|
|port|integer|The port that the cluster is listening on.|
|restorable_node_types|text[]|The list of node types that this cluster snapshot is able to restore into.|
|snapshot_create_time|timestamp without time zone|The time (in UTC format) when Amazon Redshift began the snapshot|
|snapshot_identifier|text|The snapshot identifier that is provided in the request.|
|snapshot_retention_start_time|timestamp without time zone|A timestamp representing the start of the retention period for the snapshot.|
|snapshot_type|text|The snapshot type|
|source_region|text|The source region from which the snapshot was copied.|
|status|text|The snapshot status.|
|total_backup_size_in_mega_bytes|float|The size of the complete set of backup data that would be used to restore the cluster.|
|vpc_id|text|The VPC identifier of the cluster if the snapshot is from a cluster in a VPC. Otherwise, this field is not in the output.|
|tags|jsonb|Tags consisting of a name/value pair for a resource.|
