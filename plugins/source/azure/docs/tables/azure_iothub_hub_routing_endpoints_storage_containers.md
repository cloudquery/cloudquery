
# Table: azure_iothub_hub_routing_endpoints_storage_containers
RoutingStorageContainerProperties the properties related to a storage container endpoint.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|id|text|Id of the storage container endpoint|
|connection_string|text|The connection string of the storage account.|
|endpoint_uri|text|The url of the storage endpoint|
|authentication_type|text|Method used to authenticate against the storage endpoint|
|identity_user_assigned_identity|text|The user assigned identity.|
|name|text|The name that identifies this endpoint|
|subscription_id|text|The subscription identifier of the storage account.|
|resource_group|text|The name of the resource group of the storage account.|
|container_name|text|The name of storage container in the storage account.|
|file_name_format|text|File name format for the blob|
|batch_frequency_in_seconds|integer|Time interval at which blobs are written to storage|
|max_chunk_size_in_bytes|integer|Maximum number of bytes for each blob written to storage|
|encoding|text|Encoding that is used to serialize messages to blobs|
