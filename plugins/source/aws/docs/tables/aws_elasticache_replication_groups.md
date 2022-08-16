
# Table: aws_elasticache_replication_groups
Contains all of the attributes of a specific Redis replication group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN (Amazon Resource Name) of the replication group.|
|at_rest_encryption_enabled|boolean|A flag that enables encryption at-rest when set to true|
|auth_token_enabled|boolean|A flag that enables using an AuthToken (password) when issuing Redis commands. Default: false|
|auth_token_last_modified_date|timestamp without time zone|The date the auth token was last modified|
|auto_minor_version_upgrade|boolean|Auto minor version upgrade.|
|automatic_failover|text|Indicates the status of automatic failover for this Redis replication group.|
|cache_node_type|text|The name of the compute and memory capacity node type for each node in the replication group.|
|cluster_enabled|boolean|A flag indicating whether or not this replication group is cluster enabled; i.e., whether its data can be partitioned across multiple shards (API/CLI: node groups)|
|configuration_endpoint_address|text|The DNS hostname of the cache node.|
|configuration_endpoint_port|bigint|The port number that the cache engine is listening on.|
|data_tiering|text|Enables data tiering|
|description|text|The user supplied description of the replication group.|
|global_replication_group_id|text|The name of the Global datastore|
|global_replication_group_member|text|The role of the replication group in a Global datastore|
|kms_key_id|text|The ID of the KMS key used to encrypt the disk in the cluster.|
|member_clusters|text[]|The names of all the cache clusters that are part of this replication group.|
|member_clusters_outpost_arns|text[]|The outpost ARNs of the replication group's member clusters.|
|multi_az|text|A flag indicating if you have Multi-AZ enabled to enhance fault tolerance|
|pending_auth_token_status|text|Pending modified auth token status|
|pending_automatic_failover_status|text|pending autmatic failover for this redis replication group|
|pending_primary_cluster_id|text|The primary cluster ID that is applied immediately (if --apply-immediately was specified), or during the next maintenance window.|
|pending_resharding_slot_migration_progress_percentage|float|The percentage of the slot migration that is complete.|
|pending_user_group_ids_to_add|text[]|The ID of the user group to add.|
|pending_user_group_ids_to_remove|text[]|The ID of the user group to remove.|
|replication_group_create_time|timestamp without time zone|The date and time when the cluster was created.|
|replication_group_id|text|The identifier for the replication group.|
|snapshot_retention_limit|bigint|The number of days for which ElastiCache retains automatic cluster snapshots before deleting them|
|snapshot_window|text|The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard)|
|snapshotting_cluster_id|text|The cluster ID that is used as the daily snapshot source for the replication group.|
|status|text|The current state of this replication group - creating, available, modifying, deleting, create-failed, snapshotting.|
|transit_encryption_enabled|boolean|A flag that enables in-transit encryption when set to true|
|user_group_ids|text[]|The ID of the user group associated to the replication group.|
