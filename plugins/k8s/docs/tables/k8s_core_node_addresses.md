
# Table: k8s_core_node_addresses
NodeAddress contains information for the node's address.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|node_cq_id|uuid|Unique CloudQuery ID of k8s_core_nodes table (FK)|
|type|text|Node address type, one of Hostname, ExternalIP or InternalIP.|
|address|text|The node address.|
