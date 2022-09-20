
# Table: azure_network_route_tables
Azure route table
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|id|text|Resource ID.|
|route_table_subnets|text[]|A collection of references to subnets.|
|disable_bgp_route_propagation|boolean|Whether to disable the routes learned by BGP on that route table.|
|provisioning_state|text|The provisioning state of the route table resource.|
|resource_guid|text|The resource GUID property of the route table.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|name|text|Resource name.|
|type|text|Resource type.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
