
# Table: azure_cdn_profile_endpoint_origins
DeepCreatedOrigin the main origin of CDN content which is added when creating a CDN endpoint
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|name|text|Origin name which must be unique within the endpoint|
|host_name|text|The address of the origin|
|http_port|bigint|The value of the HTTP port|
|https_port|bigint|The value of the HTTPS port|
|origin_host_header|text|The host header value sent to the origin with each request|
|priority|bigint|Priority of origin in given origin group for load balancing|
|weight|bigint|Weight of the origin in given origin group for load balancing|
|enabled|boolean|Origin is enabled for load balancing or not|
|private_link_alias|text|The Alias of the Private Link resource|
|private_link_resource_id|text|The Resource Id of the Private Link resource|
|private_link_location|text|The location of the Private Link resource|
|private_link_approval_message|text|A custom message to be included in the approval request to connect to the Private Link|
