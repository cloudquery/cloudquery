
# Table: k8s_core_node_images
List of container images on this node.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|node_cq_id|uuid|Unique CloudQuery ID of k8s_core_nodes table (FK)|
|names|text[]|Names by which this image is known.|
|size_bytes|bigint|The size of the image in bytes.|
