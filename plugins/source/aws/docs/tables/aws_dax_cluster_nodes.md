
# Table: aws_dax_cluster_nodes
Represents an individual node within a DAX cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_dax_clusters table (FK)|
|availability_zone|text|The Availability Zone (AZ) in which the node has been deployed.|
|endpoint_address|text|The DNS hostname of the endpoint.|
|endpoint_port|integer|The port number that applications should use to connect to the endpoint.|
|endpoint_url|text|The URL that applications should use to connect to the endpoint|
|node_create_time|timestamp without time zone|The date and time (in UNIX epoch format) when the node was launched.|
|node_id|text|A system-generated identifier for the node.|
|node_status|text|The current status of the node|
|parameter_group_status|text|The status of the parameter group associated with this node|
