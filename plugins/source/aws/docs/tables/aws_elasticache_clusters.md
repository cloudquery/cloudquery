
# Table: aws_elasticache_clusters
Contains all of the attributes of a specific cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN (Amazon Resource Name) of the cache cluster.|
|at_rest_encryption_enabled|boolean|A flag that enables encryption at-rest when set to true|
|auth_token_enabled|boolean|A flag that enables using an AuthToken (password) when issuing Redis commands. Default: false|
|auth_token_last_modified_date|timestamp without time zone|The date the auth token was last modified|
|auto_minor_version_upgrade|boolean|Auto minor version upgrade|
|create_time|timestamp without time zone|The date and time when the cluster was created.|
|id|text|The user-supplied identifier of the cluster|
|status|text|The current state of this cluster, one of the following values: available, creating, deleted, deleting, incompatible-network, modifying, rebooting cluster nodes, restore-failed, or snapshotting.|
|cache_node_type|text|The name of the compute and memory capacity node type for the cluster|
|cache_parameter_group_cache_node_ids_to_reboot|text[]|A list of the cache node IDs which need to be rebooted for parameter changes to be applied|
|cache_parameter_group_name|text|The name of the cache parameter group.|
|cache_parameter_group_parameter_apply_status|text|The status of parameter updates.|
|cache_subnet_group_name|text|The name of the cache subnet group associated with the cluster.|
|client_download_landing_page|text|The URL of the web page where you can download the latest ElastiCache client library.|
|configuration_endpoint_address|text|The DNS hostname of the cache node.|
|configuration_endpoint_port|bigint|The port number that the cache engine is listening on.|
|engine|text|The name of the cache engine (memcached or redis) to be used for this cluster.|
|engine_version|text|The version of the cache engine that is used in this cluster.|
|notification_configuration_topic_arn|text|The arn of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)|
|notification_configuration_topic_status|text|The current state of a notification topic used for publishing ElastiCache events to subscribers using Amazon Simple Notification Service (SNS)|
|num_cache_nodes|bigint|The number of cache nodes in the cluster|
|pending_auth_token_status|text|Auth token status that is applied to the cluster in the future or is currently being applied|
|pending_cache_node_ids_to_remove|text[]|A list of cache node IDs that are being removed (or will be removed) from the cluster|
|pending_cache_node_type|text|The cache node type that this cluster or replication group is scaled to.|
|pending_engine_version|text|Cache engine version that is being applied to the cluster (or will be applied)|
|pending_num_cache_nodes|bigint|The new number of cache nodes for the cluster|
|preferred_availability_zone|text|The name of the Availability Zone in which the cluster is located or "Multiple" if the cache nodes are located in different Availability Zones.|
|preferred_maintenance_window|text|Specifies the weekly time range during which maintenance on the cluster is performed|
|preferred_outpost_arn|text|The outpost ARN in which the cache cluster is created.|
|replication_group_id|text|The replication group to which this cluster belongs|
|replication_group_log_delivery_enabled|boolean|A boolean value indicating whether log delivery is enabled for the replication group.|
|snapshot_retention_limit|bigint|The number of days for which ElastiCache retains automatic cluster snapshots before deleting them|
|snapshot_window|text|The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your cluster|
|transit_encryption_enabled|boolean|A flag that enables in-transit encryption when set to true|
