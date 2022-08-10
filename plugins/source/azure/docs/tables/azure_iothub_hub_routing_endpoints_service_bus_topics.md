
# Table: azure_iothub_hub_routing_endpoints_service_bus_topics
RoutingServiceBusTopicEndpointProperties the properties related to service bus topic endpoint types.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|id|text|Id of the service bus topic endpoint|
|connection_string|text|The connection string of the service bus topic endpoint.|
|endpoint_uri|text|The url of the service bus topic endpoint|
|entity_path|text|Queue name on the service bus topic|
|authentication_type|text|Method used to authenticate against the service bus topic endpoint|
|identity_user_assigned_identity|text|The user assigned identity.|
|name|text|The name that identifies this endpoint|
|subscription_id|text|The subscription identifier of the service bus topic endpoint.|
|resource_group|text|The name of the resource group of the service bus topic endpoint.|
