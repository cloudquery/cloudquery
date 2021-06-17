
# Table: gcp_compute_backend_service_backends
Message containing information of one individual backend
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|backend_service_id|uuid|Unique ID of gcp_compute_backend_services table (FK)|
|balancing_mode|text|Specifies how to determine whether the backend of a load balancer can handle additional traffic or is fully loaded For usage guidelines, see  Connection balancing mode|
|capacity_scaler|float|A multiplier applied to the backend's target capacity of its balancing mode The default value is 1, which means the group serves up to 100% of its configured capacity (depending on balancingMode) A setting of 0 means the group is completely drained, offering 0% of its available capacity The valid ranges are 00 and [01,10] You cannot configure a setting larger than 0 and smaller than 01 You cannot configure a setting of 0 when there is only one backend attached to the backend service|
|description|text|An optional description of this resource Provide this property when you create the resource|
|failover|boolean|This field designates whether this is a failover backend More than one failover backend can be configured for a given BackendService|
|group|text|The fully-qualified URL of an instance group or network endpoint group (NEG) resource The type of backend that a backend service supports depends on|
|max_connections|bigint|Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode Not available if the backend's balancingMode is RATE|
|max_connections_per_endpoint|bigint|Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is RATE Not supported by:  - Internal TCP/UDP Load Balancing - Network Load Balancing|
|max_connections_per_instance|bigint|Defines a target maximum number of simultaneous connections For usage guidelines, see Connection balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is RATE Not supported by:  - Internal TCP/UDP Load Balancing - Network Load Balancing|
|max_rate|bigint|Defines a maximum number of HTTP requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION|
|max_rate_per_endpoint|float|Defines a maximum target for requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION|
|max_rate_per_instance|float|Defines a maximum target for requests per second (RPS) For usage guidelines, see Rate balancing mode and Utilization balancing mode  Not available if the backend's balancingMode is CONNECTION|
|max_utilization|float||
