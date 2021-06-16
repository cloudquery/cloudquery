
# Table: aws_redshift_cluster_nodes
The identifier of a node in a cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique ID of aws_redshift_clusters table (FK)|
|node_role|text|Whether the node is a leader node or a compute node.|
|private_ip_address|text|The private IP address of a node within a cluster.|
|public_ip_address|text|The public IP address of a node within a cluster.|
