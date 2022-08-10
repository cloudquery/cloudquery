
# Table: azure_iothub_hub_routing_endpoints_service_bus_queues
RoutingServiceBusQueueEndpointProperties the properties related to service bus queue endpoint types.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|id|text|Id of the service bus queue endpoint|
|connection_string|text|The connection string of the service bus queue endpoint.|
|endpoint_uri|text|The url of the service bus queue endpoint|
|entity_path|text|Queue name on the service bus namespace|
|authentication_type|text|Method used to authenticate against the service bus queue endpoint|
|identity_user_assigned_identity|text|The user assigned identity.|
|name|text|The name that identifies this endpoint|
|subscription_id|text|The subscription identifier of the service bus queue endpoint.|
|resource_group|text|The name of the resource group of the service bus queue endpoint.|
