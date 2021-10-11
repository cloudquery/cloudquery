
# Table: k8s_core_service_ports
The list of ports that are exposed by this service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of k8s_core_services table (FK)|
|name|text|The name of this port within the service|
|protocol|text|The IP protocol for this port|
|app_protocol|text|The application protocol for this port.|
|port|integer|The port that will be exposed by this service.|
|target_port_type|bigint|Port type, integer or string|
|target_port_int_val|integer|Port as an integer value.|
|target_port_str_val|text|Port as a string value.|
|node_port|integer|The port on each node on which this service is exposed when type is NodePort or LoadBalancer|
