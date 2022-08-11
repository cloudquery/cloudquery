
# Table: k8s_core_node_volumes_attached
List of volumes that are attached to the node.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|node_cq_id|uuid|Unique CloudQuery ID of k8s_core_nodes table (FK)|
|name|text|Name of the attached volume.|
|device_path|text|Device path where the volume should be available.|
