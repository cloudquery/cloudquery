
# Table: azure_container_managed_cluster_pip_user_assigned_identity_exceptions
ManagedClusterPodIdentityException
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_cluster_cq_id|uuid|Unique CloudQuery ID of azure_container_managed_clusters table (FK)|
|name|text|Name of the pod identity exception|
|namespace|text|Namespace of the pod identity exception|
|pod_labels|jsonb|Pod labels to match|
