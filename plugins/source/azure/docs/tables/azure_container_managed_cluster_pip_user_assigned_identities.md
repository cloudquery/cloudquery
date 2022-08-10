
# Table: azure_container_managed_cluster_pip_user_assigned_identities
ManagedClusterPodIdentity
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|managed_cluster_cq_id|uuid|Unique CloudQuery ID of azure_container_managed_clusters table (FK)|
|name|text|Name of the pod identity|
|namespace|text|Namespace of the pod identity|
|binding_selector|text|Binding selector to use for the AzureIdentityBinding resource|
|identity_resource_id|text|The resource id of the user assigned identity|
|identity_client_id|text|The client id of the user assigned identity|
|identity_object_id|text|The object id of the user assigned identity|
|provisioning_state|text|The current provisioning state of the pod identity.|
