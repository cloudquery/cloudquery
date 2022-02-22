
# Table: azure_iothub_hub_routing_endpoints_event_hubs
RoutingEventHubProperties the properties related to an event hub endpoint.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|id|text|Id of the event hub endpoint|
|connection_string|text|The connection string of the event hub endpoint.|
|endpoint_uri|text|The url of the event hub endpoint|
|entity_path|text|Event hub name on the event hub namespace|
|authentication_type|text|Method used to authenticate against the event hub endpoint|
|identity_user_assigned_identity|text|The user assigned identity.|
|name|text|The name that identifies this endpoint|
|subscription_id|text|The subscription identifier of the event hub endpoint.|
|resource_group|text|The name of the resource group of the event hub endpoint.|
