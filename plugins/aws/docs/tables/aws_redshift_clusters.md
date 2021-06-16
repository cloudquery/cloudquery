
# Table: aws_redshift_clusters
Describes a cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|allow_version_upgrade|boolean|A boolean value that, if true, indicates that major version upgrades will be applied automatically to the cluster during the maintenance window.|
|automated_snapshot_retention_period|integer|The number of days that automatic cluster snapshots are retained.|
|availability_zone|text|The name of the Availability Zone in which the cluster is located.|
|availability_zone_relocation_status|text|Describes the status of the Availability Zone relocation operation.|
|cluster_availability_status|text|The availability status of the cluster for queries.|
|cluster_create_time|timestamp without time zone|The date and time that the cluster was created.|
|cluster_identifier|text|The unique identifier of the cluster.|
|cluster_namespace_arn|text|The namespace Amazon Resource Name (ARN) of the cluster.|
|cluster_public_key|text|The public key for the cluster.|
|cluster_revision_number|text|The specific revision number of the database in the cluster.|
|cluster_snapshot_copy_status_destination_region|text|The destination region that snapshots are automatically copied to when cross-region snapshot copy is enabled.|
|cluster_snapshot_copy_status_manual_snapshot_retention_period|integer|The number of days that automated snapshots are retained in the destination region after they are copied from a source region.|
|cluster_snapshot_copy_status_retention_period|bigint|The number of days that automated snapshots are retained in the destination region after they are copied from a source region.|
|cluster_snapshot_copy_status_snapshot_copy_grant_name|text|The name of the snapshot copy grant.|
|cluster_status|text|The current state of the cluster.|
|cluster_subnet_group_name|text|The name of the subnet group that is associated with the cluster.|
|cluster_version|text|The version ID of the Amazon Redshift engine that is running on the cluster.|
|db_name|text|The name of the initial database that was created when the cluster was created.|
|data_transfer_progress_current_rate_in_mega_bytes_per_second|float|Describes the data transfer rate in MB's per second.|
|data_transfer_progress_data_transferred_in_mega_bytes|bigint|Describes the total amount of data that has been transferred in MB's.|
|data_transfer_progress_elapsed_time_in_seconds|bigint|Describes the number of seconds that have elapsed during the data transfer.|
|data_transfer_progress_estimated_time_to_completion_in_seconds|bigint|Describes the estimated number of seconds remaining to complete the transfer.|
|data_transfer_progress_status|text|Describes the status of the cluster.|
|data_transfer_progress_total_data_in_mega_bytes|bigint|Describes the total amount of data to be transferred in megabytes.|
|elastic_ip_status_elastic_ip|text|The elastic IP (EIP) address for the cluster.|
|elastic_ip_status|text|The status of the elastic IP (EIP) address.|
|elastic_resize_number_of_node_options|text|The number of nodes that you can resize the cluster to with the elastic resize method.|
|encrypted|boolean|A boolean value that, if true, indicates that data in the cluster is encrypted at rest.|
|endpoint_address|text|The DNS address of the Cluster.|
|endpoint_port|integer|The port that the database engine is listening on.|
|enhanced_vpc_routing|boolean|An option that specifies whether to create the cluster with enhanced VPC routing enabled.|
|expected_next_snapshot_schedule_time|timestamp without time zone|The date and time when the next snapshot is expected to be taken for clusters with a valid snapshot schedule and backups enabled.|
|expected_next_snapshot_schedule_time_status|text|The status of next expected snapshot for clusters having a valid snapshot schedule and backups enabled.|
|hsm_status_hsm_client_certificate_identifier|text|Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.|
|hsm_status_hsm_configuration_identifier|text|Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.|
|hsm_status|text|Reports whether the Amazon Redshift cluster has finished applying any HSM settings changes specified in a modify cluster command.|
|kms_key_id|text|The AWS Key Management Service (AWS KMS) key ID of the encryption key used to encrypt data in the cluster.|
|maintenance_track_name|text|The name of the maintenance track for the cluster.|
|manual_snapshot_retention_period|integer|The default number of days to retain a manual snapshot.|
|master_username|text|The master user name for the cluster.|
|modify_status|text|The status of a modify operation, if any, initiated for the cluster.|
|next_maintenance_window_start_time|timestamp without time zone|The date and time in UTC when system maintenance can begin.|
|node_type|text|The node type for the nodes in the cluster.|
|number_of_nodes|integer|The number of compute nodes in the cluster.|
|pending_actions|text[]|Cluster operations that are waiting to be started.|
|pending_modified_values_automated_snapshot_retention_period|integer|The pending or in-progress change of the automated snapshot retention period.|
|pending_modified_values_cluster_identifier|text|The pending or in-progress change of the new identifier for the cluster.|
|pending_modified_values_cluster_type|text|The pending or in-progress change of the cluster type.|
|pending_modified_values_cluster_version|text|The pending or in-progress change of the service version.|
|pending_modified_values_encryption_type|text|The encryption type for a cluster.|
|pending_modified_values_enhanced_vpc_routing|boolean|An option that specifies whether to create the cluster with enhanced VPC routing enabled.|
|pending_modified_values_maintenance_track_name|text|The name of the maintenance track that the cluster will change to during the next maintenance window.|
|pending_modified_values_master_user_password|text|The pending or in-progress change of the master user password for the cluster.|
|pending_modified_values_node_type|text|The pending or in-progress change of the cluster's node type.|
|pending_modified_values_number_of_nodes|integer|The pending or in-progress change of the number of nodes in the cluster.|
|pending_modified_values_publicly_accessible|boolean|The pending or in-progress change of the ability to connect to the cluster from the public network.|
|preferred_maintenance_window|text|The weekly time range, in Universal Coordinated Time (UTC), during which system maintenance can occur.|
|publicly_accessible|boolean|A boolean value that, if true, indicates that the cluster can be accessed from a public network.|
|resize_info_allow_cancel_resize|boolean|A boolean value indicating if the resize operation can be cancelled.|
|resize_info_resize_type|text|Returns the value ClassicResize.|
|restore_status_current_restore_rate_in_mega_bytes_per_second|float|The number of megabytes per second being transferred from the backup storage.|
|restore_status_elapsed_time_in_seconds|bigint|The amount of time an in-progress restore has been running, or the amount of time it took a completed restore to finish.|
|restore_status_estimated_time_to_completion_in_seconds|bigint|The estimate of the time remaining before the restore will complete.|
|restore_status_progress_in_mega_bytes|bigint|The number of megabytes that have been transferred from snapshot storage.|
|restore_status_snapshot_size_in_mega_bytes|bigint|The size of the set of snapshot data used to restore the cluster.|
|restore_status|text|The status of the restore action.|
|snapshot_schedule_identifier|text|A unique identifier for the cluster snapshot schedule.|
|snapshot_schedule_state|text|The current state of the cluster snapshot schedule.|
|tags|jsonb|The list of tags for the cluster.|
|total_storage_capacity_in_mega_bytes|bigint|The total storage capacity of the cluster in megabytes.|
|vpc_id|text|The identifier of the VPC the cluster is in, if the cluster is in a VPC.|
