
# Table: azure_front_door_backend_pool_backends
The set of backends for the backend pool
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_backend_pool_cq_id|uuid|Unique CloudQuery ID of azure_front_door_backend_pools table (FK)|
|address|text|Location of the backend (IP address or FQDN)|
|private_link_alias|text|The Alias of the Private Link resource|
|private_link_resource_id|text|The Resource ID of the Private Link resource|
|private_link_location|text|The location of the Private Link resource|
|private_endpoint_status|text|The Approval status for the connection to the Private Link|
|private_link_approval_message|text|A custom message to be included in the approval request to connect to the Private Link|
|http_port|integer|The HTTP TCP port number|
|https_port|integer|The HTTPS TCP port number|
|enabled_state|text|Whether the use of the backend is enabled|
|priority|integer|Priority to use for load balancing|
|weight|integer|Weight of the endpoint for load balancing purposes|
|host_header|text|The value to use as the host header sent to the backend|
