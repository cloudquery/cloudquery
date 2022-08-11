
# Table: aws_dax_clusters
Information about a DAX cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The tags associated with the cluster.|
|active_nodes|integer|The number of nodes in the cluster that are active (i.e., capable of serving requests).|
|arn|text|The Amazon Resource Name (ARN) that uniquely identifies the cluster.|
|cluster_discovery_endpoint_address|text|The DNS hostname of the endpoint.|
|cluster_discovery_endpoint_port|integer|The port number that applications should use to connect to the endpoint.|
|cluster_discovery_endpoint_url|text|The URL that applications should use to connect to the endpoint|
|cluster_endpoint_encryption_type|text|The type of encryption supported by the cluster's endpoint|
|name|text|The name of the DAX cluster.|
|description|text|The description of the cluster.|
|iam_role_arn|text|A valid Amazon Resource Name (ARN) that identifies an IAM role|
|node_ids_to_remove|text[]|A list of nodes to be removed from the cluster.|
|node_type|text|The node type for the nodes in the cluster|
|notification_configuration_topic_arn|text|The Amazon Resource Name (ARN) that identifies the topic.|
|notification_configuration_topic_status|text|The current state of the topic|
|node_ids_to_reboot|text[]|The node IDs of one or more nodes to be rebooted.|
|parameter_apply_status|text|The status of parameter updates.|
|parameter_group_name|text|The name of the parameter group.|
|preferred_maintenance_window|text|A range of time when maintenance of DAX cluster software will be performed|
|sse_description_status|text|The current state of server-side encryption:  * ENABLING - Server-side encryption is being enabled.  * ENABLED - Server-side encryption is enabled.  * DISABLING - Server-side encryption is being disabled.  * DISABLED - Server-side encryption is disabled.|
|security_groups|jsonb|A list of security groups, and the status of each, for the nodes in the cluster.|
|status|text|The current status of the cluster.|
|subnet_group|text|The subnet group where the DAX cluster is running.|
|total_nodes|integer|The total number of nodes in the cluster.|
